package main

import (
	"errors"
	"fmt"
)

type DataNode struct {
	cmd     string
	desc    string
	handler func() int
}

func ShowAllCmd() {
	for i := range head {
		fmt.Println(head[i].cmd)
	}
}

func FindCmd(cmd string) (*DataNode, error) {
	if cmd == "" {
		return nil, errors.New("命令不能为空")
	}
	for i := range head {
		if head[i].cmd == cmd {
			return &head[i], nil
		}
	}
	return nil, errors.New("命令不存在")
}
