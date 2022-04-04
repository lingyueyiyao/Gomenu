package main

import (
	"fmt"
	"unsafe"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN    = 1024
	CMD_NUM     = 10
)

type DataNode struct {
	pNext   *tLinkTableNode
	cmd     string
	desc    string
	handler func() int
}

// func(a ...interface{}) interface{}

type tDataNode = DataNode

func FindCmd(head *tLinkTable, cmd string) *tDataNode {
	var pNode *tDataNode = (*tDataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	for pNode != nil {
		if pNode.cmd == cmd {
			return pNode
		}
		pNode = (*tDataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*tLinkTableNode)(unsafe.Pointer(pNode)))))
	}
	return nil
}

func ShowAllCmd(head *tLinkTable) int {
	var pNode *tDataNode = (*tDataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	for pNode != nil {
		fmt.Printf("%s - %s\n", pNode.cmd, pNode.desc)
		pNode = (*tDataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*tLinkTableNode)(unsafe.Pointer(pNode)))))
	}
	return 0
}

func InitMenuData(ppLinktable **tLinkTable) int {
	*ppLinktable = CreatLinkTable()
	pNode := new(tDataNode)
	pNode.cmd = "help"
	pNode.desc = "Menu List"
	pNode.handler = Help
	AddLinkTableNode(*ppLinktable, (*tLinkTableNode)(unsafe.Pointer(pNode)))
	pNode = new(tDataNode)
	pNode.cmd = "version"
	pNode.desc = "Menu Program v1.0"
	pNode.handler = version
	AddLinkTableNode(*ppLinktable, (*tLinkTableNode)(unsafe.Pointer(pNode)))
	pNode = new(tDataNode)
	pNode.cmd = "quit"
	pNode.desc = "Quit from Menu Program v1.0"
	pNode.handler = Quit
	AddLinkTableNode(*ppLinktable, (*tLinkTableNode)(unsafe.Pointer(pNode)))

	return 0
}

var head *tLinkTable = nil

func init() {
	InitMenuData(&head)
}

func main() {

	var cmd string
	for {
		fmt.Print("Input a cmd number > ")
		fmt.Scanln(&cmd)
		p := FindCmd(head, cmd)
		if p == nil {
			fmt.Println("This is a wrong cmd!")
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

func Help() int {
	ShowAllCmd(head)
	return 0
}

func version() int {
	fmt.Println("menu version 1.0")
	return 0
}

func Quit() int {
	fmt.Println("Menu exit")
	return -1
}
