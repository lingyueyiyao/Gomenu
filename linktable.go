package main

import (
	"sync"
)

const (
	SUCCESS = 0
	FAILURE = -1
)

type LinkTableNode struct {
	pNext *LinkTableNode
}

type tLinkTableNode = LinkTableNode

type LinkTable struct {
	pHead     *tLinkTableNode
	pTail     *tLinkTableNode
	SumOfNode int
	mutex     *sync.RWMutex
}

type tLinkTable = LinkTable

func CreatLinkTable() *tLinkTable {
	pLinkTable := new(tLinkTable)
	if pLinkTable == nil {
		return nil
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	pLinkTable.mutex = new(sync.RWMutex)
	return pLinkTable
}

func DeleteLinkTable(pLinkTable *tLinkTable) int {
	if pLinkTable == nil {
		return FAILURE
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	pLinkTable.mutex = nil
	pLinkTable = nil
	return SUCCESS
}

func AddLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return FAILURE
	}
	pNode.pNext = nil
	pLinkTable.mutex.Lock()
	defer pLinkTable.mutex.Unlock()
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
	}
	if pLinkTable.pTail == nil {
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.SumOfNode += 1
	return SUCCESS
}

func DelLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return FAILURE
	}
	pLinkTable.mutex.Lock()
	defer pLinkTable.mutex.Unlock()
	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.pNext
		pLinkTable.SumOfNode -= 1
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return SUCCESS
	}
	pTempNode := pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode.pNext == pNode {
			pTempNode.pNext = pTempNode.pNext.pNext
			pLinkTable.SumOfNode -= 1
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			return SUCCESS
		}
		pTempNode = pTempNode.pNext
	}
	return FAILURE
}

func SearchLinkTableNode(pLinkTable *tLinkTable,
	Condition func(pNode *tLinkTableNode, args interface{}) int, args interface{}) *tLinkTableNode {
	if pLinkTable == nil || Condition == nil {
		return nil
	}
	pNode := pLinkTable.pHead
	for pNode != nil {
		if Condition(pNode, args) == SUCCESS {
			return pNode
		}
		pNode = pNode.pNext
	}
	return nil
}

func GetLinkTableHead(pLinkTable *tLinkTable) *tLinkTableNode {
	if pLinkTable == nil {
		return nil
	}
	return pLinkTable.pHead
}

func GetNextLinkTableNode(pLinkTable *tLinkTable, pNode *tLinkTableNode) *tLinkTableNode {
	if pLinkTable == nil || pNode == nil {
		return nil
	}
	pTempNode := pLinkTable.pHead
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.pNext
		}
		pTempNode = pTempNode.pNext
	}
	return nil
}
