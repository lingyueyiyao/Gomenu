package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"unsafe"
)

type Node struct {
	Head tLinkTableNode
	Data int
}

type tNode = Node

func TestTable(t *testing.T) {
	pLinkTable := CreatLinkTable()
	if pLinkTable == nil {
		fmt.Println("CreateLinkTable Error!")
		os.Exit(0)
	}
	for i := 0; i < 10; i++ {
		pNode := new(tNode)
		pNode.Data = i
		log.Println("AddLinkTableNode")
		AddLinkTableNode(pLinkTable, (*tLinkTableNode)(unsafe.Pointer(pNode)))
	}
	log.Println("SearchLinktableNode")
	num := 6
	pTempNode := (*tNode)(unsafe.Pointer(SearchLinkTableNode(pLinkTable, searchCondition, num)))
	fmt.Println(pTempNode.Data)

	pTempNode = Search(pLinkTable)
	fmt.Println(pTempNode.Data)
	log.Println("DelLinkTableNode")
	DelLinkTableNode(pLinkTable, (*tLinkTableNode)(unsafe.Pointer(pTempNode)))
	pTempNode = nil
	DeleteLinkTable(pLinkTable)
}

func Search(pLinkTable *tLinkTable) *tNode {
	log.Println("Search GetLinkTableHead")
	pNode := (*tNode)(unsafe.Pointer(GetLinkTableHead(pLinkTable)))
	for pNode != nil {
		if pNode.Data == 5 {
			return pNode
		}
		log.Println("GetNextLinkTableNode")
		pNode = (*tNode)(unsafe.Pointer(GetNextLinkTableNode(pLinkTable,
			(*tLinkTableNode)(unsafe.Pointer(pNode)))))
	}
	return nil
}

func searchCondition(pLinkTableNode *tLinkTableNode, args interface{}) int {
	pNode := (*tNode)(unsafe.Pointer(pLinkTableNode))
	if pNode.Data == args {
		return SUCCESS
	}
	return FAILURE
}
