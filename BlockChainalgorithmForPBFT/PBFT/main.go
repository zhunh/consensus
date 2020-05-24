package main

import (
	"BlockChainalgorithmForPBFT/Peers"
	"fmt"
	"net/http"

	"os"
)


func main() {
	//接受终端传来的参数
	userId:=os.Args[1]
	fmt.Println(userId)
	//初始化国家的id对应的地址
	Peers.PeersTable= map[string]string{
		"N0":"localhost:1111",
		"N1":"localhost:1112",
		"N2":"localhost:1113",
		"N3":"localhost:1114",
	}

	//创建国家初始化节点id和路径
	peer:=Peers.InitPeers(userId,Peers.PeersTable[userId])
	//第一阶段向主服务器请求
    http.HandleFunc("/req",peer.Request)
	//第二阶段主服务器向其余节点进行分发
	http.HandleFunc("/prePrepare",peer.PrePrepare)

	//第三阶段节点之间相互交互
	http.HandleFunc("/prepare",peer.Prepare)
	//第四阶段开始验证
	http.HandleFunc("/commit",peer.Commit)
	err:=http.ListenAndServe(peer.Path,nil)
	if err!=nil{
		fmt.Println(err)
	}
}