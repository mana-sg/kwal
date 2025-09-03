package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
)

const (
	pidFile = "/tmp/kls.pid"
	Reset   = "\033[0m"
	Red     = "\033[31m"
	White   = "\033[37m"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		fmt.Println("Usage: kls [start|stop|restart|status]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		startServer()
	case "stop":
		stopServer()
	case "restart":
		stopServer()
		startServer()
	case "status":
		checkStatus()
	default:
		fmt.Println("start: \tStarts the API server.")
		fmt.Println("stop: \tStops the API server if running.")
		fmt.Println("restart: \tRestarts the API server.")
		fmt.Println("status: \tCheck the status of the server.")
	}
}


func startServer() {
    if data, err := os.ReadFile(pidFile); err == nil {
        if pid, _ := strconv.Atoi(string(data)); pid > 0 {
            if syscall.Kill(pid, 0) == nil {
                fmt.Println(Red, "Server already running.", Reset)
                return
            }
            _ = os.Remove(pidFile)
        }
    }

    home, _ := os.UserHomeDir()
    serverPath := filepath.Join(home, os.Getenv("SERVER_LOCATION"), "server.go")
    binaryPath := filepath.Join(os.TempDir(), "kv-server")

    buildCmd := exec.Command("go", "build", "-o", binaryPath, serverPath)
    buildCmd.Stdout = os.Stdout
    buildCmd.Stderr = os.Stderr
    if err := buildCmd.Run(); err != nil {
        fmt.Println(Red, "Error building server:", err, Reset)
        return
    }

    cmd := exec.Command(binaryPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

    if err := cmd.Start(); err != nil {
        fmt.Println(Red, "Error starting server:", err.Error(), Reset)
        return
    }

    _ = os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
    fmt.Println(White, "Server started with PID:", cmd.Process.Pid, Reset)
}

func checkStatus() {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		fmt.Println(White, "KLS daemon is not running.", Reset)
		return
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Println(Red, "Invalid PID in pid file.", Reset)
		return
	}

	_, err = os.FindProcess(pid)
	if err != nil {
		fmt.Println(Red, "Error finding process:", err.Error(), Reset)
		_ = os.Remove(pidFile)
		return
	}

	print(White, "Daemon up and running with PID: ", pid, Reset)
}


func stopServer() {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		fmt.Println(Red, "No running server found.", Reset)
		return
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Println(Red, "Invalid PID in pid file.", Reset)
		return
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(Red, "Error finding process:", err.Error(), Reset)
		_ = os.Remove(pidFile)
		return
	}

	if err := proc.Kill(); err != nil {
		fmt.Println(Red, "Error stopping server:", err.Error(), Reset)
		return
	}

	_ = os.Remove(pidFile)
	fmt.Println(White, "Server stopped.", Reset)
}
