package Block

import (
	"BlockChainalgorithmForPos/Peers"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	PreHash string
	HashCode string
	TimeStamp string
	Index int
	//区块验证者
	Validator *Peers.Peers
	Data string
}
//初始化创世块
func GenerateFirstBlock(data string,validator *Peers.Peers)Block{
	var firstblock Block
	firstblock.TimeStamp=time.Now().String()
	firstblock.Index=1
	firstblock.PreHash="0"
	firstblock.Data=data
	firstblock.Validator=validator
	firstblock.HashCode=GenerationHashValue(&firstblock)
	return firstblock
}


//计算区块的hash值

func GenerationHashValue(block *Block)string{
	var hash=block.PreHash+
		block.Data+
		strconv.Itoa(block.Index)+
		block.TimeStamp+block.Validator.Address
	sha:=sha256.New()
	sha.Write([]byte(hash))
	hashed:=sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//产生新的区块

func GenerateNextBlock(oldBlock Block,data string,validator *Peers.Peers)Block{
	var newBlock Block
	newBlock.TimeStamp=time.Now().String()
	newBlock.Index=oldBlock.Index+1
	newBlock.PreHash=oldBlock.HashCode
	newBlock.Data=data
	newBlock.Validator=validator
	//计算hash
	newBlock.HashCode=GenerationHashValue(&newBlock)
	return newBlock
}

