package main

import (
	"fmt"
	"log"
	"testing"
)

var results = [...]int{1, 1, 1}

var info = [...]string{
	"test report",
	"TC1 CreateLinkTable",
	"TC2.1...",
}

func TestTable1(t *testing.T) {
	p := CreatLinkTable()
	if p == nil {
		log.Println("TC1 fail")
		results[1] = 1
	}
	var pNode *tLinkTableNode
	ret := AddLinkTableNode(p, pNode)
	if ret == FAILURE {
		log.Println("TC2.1 Succ")
		results[2] = 1
	}

	fmt.Println("test report")
	for i := 1; i < len(results); i++ {
		if results[i] == 1 {
			fmt.Printf("Testcase Number%d F - %s\n", i, info[i])
		}
	}
}
