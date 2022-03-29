package main

import (
	"fmt"
)

var head = []DataNode{
	{"help", "Display all commands", nil},
	{"version", "menu program v1.0", version},
	{"q", "exit menu", q},
}

func init() {
	head[0].handler = help
}

func main() {

	var cmd string
	for {
		fmt.Print("请输入命令：")
		fmt.Scanln(&cmd)
		p, err := FindCmd(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s - %s\n", p.cmd, p.desc)
		if p.handler != nil {
			if p.handler() < 0 {
				return
			}
		}
	}
}

func help() int {
	ShowAllCmd()
	return 0
}

func version() int {
	fmt.Println("menu version 1.0")
	return 0
}

func q() int {
	fmt.Println("menu exit")
	return -1
}
