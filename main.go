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
		fmt.Println("Usage: kls [start|stop|restart]")
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
	default:
		fmt.Println("start: \tStarts the API server.")
		fmt.Println("stop: \tStops the API server if running.")
		fmt.Println("restart: \tRestarts the API server.")
	}
}

func startServer() {
	if data, err := os.ReadFile(pidFile); err == nil {
		if pid, _ := strconv.Atoi(string(data)); pid > 0 {
			if proc, _ := os.FindProcess(pid); proc != nil && proc.Signal(syscall.Signal(0)) == nil {
				fmt.Println(Red, "Server already running.", Reset)
				return
			}
			os.Remove(pidFile)
		}
	}
	home, _ := os.UserHomeDir()
	serverPath := filepath.Join(home, os.Getenv("SERVER_LOCATION"), "server.go")
	cmd := exec.Command("go", "run", serverPath, "&")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		fmt.Println(Red, "Error starting server:", err.Error(), Reset)
		return
	}
	os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	fmt.Println(White, "Server started with PID:", cmd.Process.Pid, Reset)
}

func stopServer() {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		fmt.Println(Red, "No running server found.", Reset)
		return
	}
	pid, _ := strconv.Atoi(string(data))
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(Red, "Error finding process:", err.Error(), Reset)
		return
	}
	if err := proc.Kill(); err != nil {
		fmt.Println(Red, "Error stopping server:", err.Error(), Reset)
		return
	}
	os.Remove(pidFile)
	fmt.Println(White, "Server stopped.", Reset)
}
