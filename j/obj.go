package main

import (
	"encoding/hex"
	"fmt"
	"crypto/sha256"
	"strings"
	"time"
	"math/rand"
)

const DT = "00"

//区块结构
type Block struct{
	preHash string
	nonce int
	timestamp string
	hash string
	data string
}
//矿工结构
type Worker struct {
	workId int

}
//挖矿函数
func mining()  {
	//1.创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	for {
		//2.输入数据
		noce := rand.Int63n(100000000)   //以一个随机数为输入数据
		fmt.Println(noce)
		hash.Write([]byte(string(noce)))
		//3.计算哈希值
		s := hash.Sum(nil)
		//4.将字符串编码为16进制格式,返回字符串
		hashcode := hex.EncodeToString(s)
		//对第一次求得的hash进行二次hash计算
		hash.Write([]byte(hashcode))
		s1 := hash.Sum(nil)
		hashcode1 := hex.EncodeToString(s1)
		fmt.Println(hashcode1)
		if strings.HasPrefix(hashcode1,DT){
			fmt.Println("挖矿成功：",hashcode1,"\n时间：",time.Now())
			break
		}
		//time.Sleep(200*time.Millisecond)
	}
}

func main()  {
	blockchain := make([]Block,5)
	mining()
	fmt.Println(len(blockchain))
}