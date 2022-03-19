package main

import (
	"fmt"
	"os"
)

func main() {

	var cmd string
	for {
		fmt.Print("请输入命令：")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("This is help cmd!")
		} else if cmd == "quit" {
			fmt.Println("Menu已退出!")
			os.Exit(0)
		} else {
			fmt.Println("Wrong cmd!")
		}
	}
}
