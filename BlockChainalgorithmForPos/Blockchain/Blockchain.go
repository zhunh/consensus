package Blockchain

import (
	"BlockChainalgorithmForPos/Block"
	"fmt"
)

//用linklist实现区块链链


type Node struct {
	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block

}
//指针始终指向最后一个节点为了方便添加

var LastNode *Node
// 创建头结点，保存创世区块
func CreateHeaderNode(data *Block.Block)*Node{
	//先去初始化
	var headerNode=new(Node)
	//指针域指向nil
	headerNode.NextNode=nil
	//数据域
	headerNode.Data=data
	//返回头结点
	LastNode=headerNode
	return headerNode
}

//当挖矿成功的时候，添加区块，添加节点
func AddNode(data *Block.Block)*Node{
	//创建新的节点
	var newNode=new(Node)
	newNode.NextNode=nil
	newNode.Data=data
	LastNode.NextNode=newNode
	LastNode=newNode
	return newNode
}
//查看linklist的数据
func ShowNodes(node *Node){
	n:=node
	for  {
		if n.NextNode==nil{
			fmt.Println(n.Data)
			break
		}else{
			fmt.Println(n.Data)
			n=n.NextNode
		}

	}
}