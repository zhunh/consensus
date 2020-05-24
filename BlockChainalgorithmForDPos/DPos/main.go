package main

import (
	"BlockChainalgorithmForDPos/Block"
	"BlockChainalgorithmForDPos/Blockchain"

	"fmt"
	"math/rand"
	"time"
)

//创建10个全节点
var peers=make([]Block.Peers,10)
var sortpeers=make([]Block.Peers,10)


//初始化记账节点
func InitPeers()  {
	for i:=0;i<10;i++ {
		name :=fmt.Sprintf("节点%d,我的票数为",i)
		peers[i]=Block.Peers{name,0}
	}

}
//模拟投票
func Vote()  {
	var ik int
	for i:=0;i<10;i++{

		rand.Seed(time.Now().UnixNano())
		time.Sleep(10000000)
		vote:=rand.Intn(1000)
		ik+=vote
		//用随机数为10个节点投票
		peers[i].Votes=vote
		fmt.Printf("节点[%d],票数[%d]\n",i,vote)
	}
	fmt.Println("总票数：",ik)
}
//此方法将排名选取票数前三多的节点
//采用快速排序法
func Chose()[]Block.Peers{
	copy(sortpeers,peers)
	//fmt.Println("数组")
	//fmt.Println(sortpeers)
	Quicksort(0,len(sortpeers)-1)
   return sortpeers
}
func Quicksort(left int, right int) {

	if left >= right { //需要排序的起始位置大于或等于终止位置，就表明不再需要排序
		return
	}

	i := left                     //左边的游标
	j := right                    //右边的游标
	temp := sortpeers[left].Votes // 基数

	for i != j {

		for sortpeers[j].Votes >= temp && i < j { //直到遇到了 <temp 的数就停下来
			j--

		}

		for sortpeers[i].Votes <= temp && i < j { //直到遇到了 >temp 的数就停下来
			i++
		}

		//交换这两个数
		if i < j {
			sortpeers[i], sortpeers[j] = sortpeers[j], sortpeers[i]
		}

	}

	//将基数归位
	sortpeers[i].Votes, sortpeers[left].Votes = temp, sortpeers[i].Votes
	//递归处理基数左边未处理的
	Quicksort(left, i-1)

	//递归处理基数右边未处理的
	if i != len(sortpeers)-1 {
		Quicksort(i+1, right)
	}
}

func main(){
	//初始化10个全节点
	InitPeers()
	fmt.Println(peers)

	//投票
	Vote()
	fmt.Println()
 	//排序并选取其三的节点
 	var tmp=Chose()[len(sortpeers)-3:len(sortpeers)]
 	fmt.Println(tmp)

 	//初始化区块与区块链
 	firstblock:=Block.GenerateFirstBlock("创世块",&tmp[2])

 	firstnode:=Blockchain.CreateHeaderNode(&firstblock)
	oldblcok:=firstblock

 	//由票数最高的三个节点来记账
	for i:=len(tmp)-1;i>=0; i--{

		curblock:=tmp[i].GenerateNextBlock(oldblcok,"新的区块")
		oldblcok=curblock
		Blockchain.AddNode(&curblock)

	}

	//显示区块链的数据
	Blockchain.ShowNodes(firstnode)



}