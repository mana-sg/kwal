package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mana-sg/kv-log-store/pkg/storage"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	White = "\033[37m"
)

func main() {
	storage.Kv_store.BuildStore()

	fmt.Println("\003[H\033[2J")
	fmt.Println("Welcome to my kv store!")

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
			}
			err := storage.Kv_store.Remove(objects[1])
			if err != nil {
				fmt.Println(Red, "err removing value: \n", err, Reset)
			}
			fmt.Println("key: ", objects[1], ", removed succesfully\n")
		}
	}
}
