package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"

	"github.com/mana-sg/kv-log-store/pkg/storage"
	"github.com/mana-sg/kv-log-store/pkg/wal"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	White = "\033[37m"
)

func main() {
	storage.Kv_store.BuildStore()

	fmt.Print("\033[H\033[2J")
	fmt.Println(White, "Welcome to my kv store!", Reset)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		objects := strings.Split(scanner.Text(), " ")

		switch objects[0] {
		case "set":
			if len(objects) != 3 {
				fmt.Println(Red, "Expected 2 arguments received: ", len(objects)-1, Reset)
				fmt.Println(White, "Usage:\n\tset <key> <value>", Reset)
				break
			}
			err := storage.Kv_store.Set(objects[1], objects[2])
			if err != nil {
				fmt.Println(Red, "error setting value: \n", err, Reset)
			}
			fmt.Println("Key-Value pair added succesfully\n")

		case "get":
			if len(objects) != 2 {
				fmt.Println(Red, "Expected 1 arguments received: ", len(objects)-1, Reset)
				fmt.Println(White, "Usage:\n\tget <key>", Reset)
				break
			}
			val, err := storage.Kv_store.Get(objects[1])
			if err != nil {
				fmt.Println(Red, "err getting value: \n", err, Reset)
			}
			fmt.Println(val, "\n")

		case "del":
			if len(objects) != 2 {
				fmt.Println(Red, "Expected 1 arguments received: ", len(objects)-1, Reset)
				fmt.Println(White, "Usage:\n\tdel <key>", Reset)
				break
			}
			err := storage.Kv_store.Remove(objects[1])
			if err != nil {
				fmt.Println(Red, "err removing value: \n", err, Reset)
			}
			fmt.Println("key: ", objects[1], ", removed succesfully\n")

		case "exit":
			os.Exit(0)

		case "cls":
			fmt.Print("\033[H\033[2J")
			fmt.Println(White, "Welcome to my kv store!", Reset)

		case "size":
			home, err := os.UserHomeDir() 			
			filePath := home + "/kls/log.bin"
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				log.Fatal(err)
			}

			fileSize := fileInfo.Size()
			fmt.Println(White, "File size: ", fileSize,"bytes\n", Reset)

		case "compact":
			savings, err := wal.Compact()
			if err != nil {
				fmt.Println(Red, "error compacting: %v", err, Reset)
			}

			fmt.Println(White, "\nFile comaction succesful.", )
			fmt.Println("File sie decreased by: ", savings * 100, "%\n", Reset)

		case "help":
			fallthrough

		default:
			fmt.Println(White, "\nset: \tSets a key value pair.")
			fmt.Println("\tUsage: set <key> <value>\n")

			fmt.Println("get: \tGets the value for a given key.")
			fmt.Println("\tUsage: get <key>\n")

			fmt.Println("del: \tDeletes a key value pair from the database.")
			fmt.Println("\tUsage: del <key>\n")

			fmt.Println("size: \tGets the size of the log file.")
			fmt.Println("\tUsage: size\n")

			fmt.Println("cls: \tClears the screen")
			fmt.Println("\tUsage: del <key>\n", Reset)
		}
	}
}
