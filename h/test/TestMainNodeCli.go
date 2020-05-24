package main

import (
	"consensus/node"
	"fmt"
	"os"
)
func main() {
	var a node.Node
	a.Port = os.Args[1]
	a.Token = os.Args[2]
	//a.Ip = "127.0.0.1"
	//if os.Args[3] != ""{
	//	a.Ip = os.Args[3]
	//}
	a.Ledger = "default"

	fmt.Println(a)
	cli := a.ConnectTo("192.168.0.100:10010")
	go a.Register(cli)
	//作为服务端开始监听请求
	go a.AsServer()
	for{
		;
	}

}
