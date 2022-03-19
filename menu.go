package main

import (
	"fmt"
	"os"
)

func main() {

	var cmd string
	for {
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("This is help cmd!")
		} else if cmd == "quit" {
			os.Exit(0)
		} else {
			fmt.Println("Wrong cmd!")
		}
	}
}
