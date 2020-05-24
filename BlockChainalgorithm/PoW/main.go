package main

import (
	"BlockChainalgorithm/Block"
	"BlockChainalgorithm/Blockchain"
)

func main() {
	var first = Block.GenerateFirstBlock("创世区块")
	var second=Block.GenerateNextBlock("第二块区块",first)

	//创建链表
	header:=Blockchain.CreateHeaderNode(&first)
	Blockchain.AddNode(&second)

	Blockchain.ShowNodes(header)
}