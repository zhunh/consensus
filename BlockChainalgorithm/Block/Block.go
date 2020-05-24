package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	PreHash string
	HashCode string
	TimeStamp string
	Diff int
	Data string
	Index int
	Nouce int
}
//初始化创世块
func GenerateFirstBlock(data string)Block{
	var firstblock Block
	firstblock.Data=data
	firstblock.Diff=4
	firstblock.TimeStamp=time.Now().String()
	firstblock.Index=1
	firstblock.Nouce=0
	firstblock.PreHash="0"
	firstblock.HashCode=GenerationHashValue(firstblock)
	return firstblock
}


//计算区块的hash值

func GenerationHashValue(block Block)string{
	var hash=block.PreHash+block.Data+strconv.Itoa(block.Nouce)+strconv.Itoa(block.Diff)+strconv.Itoa(block.Index)+block.TimeStamp
	sha:=sha256.New()
	sha.Write([]byte(hash))
	hashed:=sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//产生新的区块

func GenerateNextBlock(data string,oldBlock Block)Block{
	var newBlock Block
	newBlock.TimeStamp=time.Now().String()
	newBlock.Index=oldBlock.Index+1
	newBlock.Data=data
	newBlock.PreHash=oldBlock.HashCode
	newBlock.Diff=4
	newBlock.Nouce=0
	//实现pow
	newBlock.HashCode=Pow(newBlock.Diff,&newBlock)
	return newBlock
}
//挖矿
func Pow(diff int,block *Block) string{
	for  {
		hash:=GenerationHashValue(*block)
		fmt.Println(hash)
		if strings.HasPrefix(hash,strings.Repeat("0",diff)){
			fmt.Println("success")
			return hash
		}else{
			block.Nouce++
		}
	}
}