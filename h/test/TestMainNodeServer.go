package main

import (
	"consensus/node"
	"fmt"
)
func main() {
	s := make([]node.Node,1)
	var a node.MainNode = node.MainNode{
		"10010",
		s,
	}
	fmt.Println(a)


	//这个go程怎么和主go程通信（共享一个数据）
	go a.AsServer()
	var cmd string
	var n node.Node
	for{

		fmt.Scanln(&cmd)
		if cmd == "st"{
			a.StartMining()
		}
		if cmd == "pos"{
			n = a.Pos()
			fmt.Println("Pos共识获得记账权的节点为",n)
		}
	}
}
