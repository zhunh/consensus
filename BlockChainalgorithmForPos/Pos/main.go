package main

import (
	"BlockChainalgorithmForPos/Block"
	"BlockChainalgorithmForPos/Blockchain"
	"BlockChainalgorithmForPos/Peers"
	"fmt"
	"math/rand"
	"time"
)

//创建五个全节点
var peers=make([]Peers.Peers,5)
//存放参与记账的节点，后面依靠该数组去争取挖矿资格
var address=make([]*Peers.Peers,40)

//初始化记账节点
func InitPeers()  {

	peers[0]=Peers.Peers{1,1,"0x111"}
	peers[1]=Peers.Peers{2,1,"0x222"}
	peers[2]=Peers.Peers{3,1,"0x333"}
	peers[3]=Peers.Peers{4,1,"0x444"}
	peers[4]=Peers.Peers{5,6,"0x555"}
	//概率分布的下标
	count:=0
	//根据持币量乘以持币时间将节点地址按数量分散到address数组，并随机抽取
	for i:=0;i<len(peers);i++{
		for j:=0;j<peers[i].Token*peers[i].Days ;j++  {
			address[count]=&peers[i]
			count++
		}
	}

}
func Pos(oldBlock Block.Block,data string) Block.Block {
	time.Sleep(10000)
	//设计实时时间为随机种子，保证了不是伪随机
	rand.Seed(time.Now().Unix())
	//从adrress的中随机节点出来进行记账
	var rd=rand.Intn(len(address))
	peernode:=address[rd]
	fmt.Println("获得记账权的节点为：",peernode)
	newblock:=Block.GenerateNextBlock(oldBlock,data ,peernode)

 	return newblock
}
func main() {
	InitPeers()
	//创世块
	var first = Block.GenerateFirstBlock("创世块",&Peers.Peers{2,1,"0x222"})
	//
	//
	////创建区块链表
	header:=Blockchain.CreateHeaderNode(&first)
	second:=Pos(first,"第二个区块")
	Blockchain.AddNode(&second)

	Blockchain.ShowNodes(header)
}