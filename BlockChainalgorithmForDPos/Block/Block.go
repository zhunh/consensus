package Block

import (

	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)
type Peers struct {
	//节点名称
	Name string
	//被投票的个数
	Votes int
}
type Block struct {
	PreHash string
	HashCode string
	TimeStamp string
	Index int
	//区块验证者
	Delegate *Peers
	Data string
}
//初始化创世块
func GenerateFirstBlock(data string,delegate *Peers)Block{
	var firstblock Block
	firstblock.TimeStamp=time.Now().String()
	firstblock.Index=1
	firstblock.PreHash="0"
	firstblock.Data=data
	firstblock.Delegate=delegate
	firstblock.HashCode=GenerationHashValue(&firstblock)
	return firstblock
}


//计算区块的hash值

func GenerationHashValue(block *Block)string{
	var hash=block.PreHash+
		block.Data+
		strconv.Itoa(block.Index)+
		block.TimeStamp
	sha:=sha256.New()
	sha.Write([]byte(hash))
	hashed:=sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//产生新的区块



//使用被选举的节点来创建新的区块
func (peer *Peers) GenerateNextBlock(oldBlock Block,data string)Block{
	var newBlock Block
	newBlock.TimeStamp=time.Now().String()
	newBlock.Index=oldBlock.Index+1
	newBlock.PreHash=oldBlock.HashCode
	newBlock.Data=data
	newBlock.Delegate=peer
	//计算hash
	newBlock.HashCode=GenerationHashValue(&newBlock)
	return newBlock
}