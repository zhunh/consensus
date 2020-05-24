package node

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	pd "consensus/grpc"
	"context"
	"strings"
	"time"
)

type Node struct {
	Ip  string
	Port string
	Ledger string
	Token string
}

func (c *Node) PowBroadcast(ctx context.Context, in *pd.PowReqMsg) (*pd.PowRspMsg, error) {
	return &pd.PowRspMsg{PowRspMsg:"hello"},nil
}
func (c *Node) PosBroadcast(ctx context.Context, in *pd.PosReqMsg) (*pd.PosRspMsg, error) {
	return &pd.PosRspMsg{PosRspMsg:"good"},nil
}
func (c *Node) StartMining(ctx context.Context, in *pd.StartMiningReq) (*pd.StartMiningRsp, error) {
	c.Mining(in.DT)
	return &pd.StartMiningRsp{StartMiningRsp:"已开始挖矿"}, nil
}

func (c *Node)AsServer()  {
	fmt.Println("服务准备启动",c.Port)
	//创建网络
	ln ,err :=net.Listen("tcp",":"+c.Port)
	if err !=nil{
		fmt.Println("网络错误",err)
	}

	//创建grpc的服务
	srv:=grpc.NewServer()

	//注册服务
	pd.RegisterNodeServerServer(srv,c)

	fmt.Println("服务已启动",c.Port)

	//等待网络连接
	err=srv.Serve(ln)
	if err!=nil {
		fmt.Println("网络错误",err)
	}

}
//addr  "127.0.0.1:10086"
func (c *Node)ConnectTo(addr string) pd.MainNodeServiceClient {
	fmt.Printf("%s开始连接%s\n",c.Port,addr)
	//客户端连接服务器
	conn ,err :=grpc.Dial(addr,grpc.WithInsecure())
	if err!=nil {
		fmt.Println("网络异常",err)
	}
	//网络延迟关闭
	//defer  conn.Close()
	//获得grpc句柄
	cli:=pd.NewMainNodeServiceClient(conn)
	fmt.Println(c.Port+"已连接到"+addr)
	return cli
}

func (c *Node)Invoke(cli pd.NodeServerClient)  {
	//通过句柄调用函数
	re ,err :=cli.PowBroadcast(context.Background(),&pd.PowReqMsg{PowReqMsg:"Kristin"})
	if err!=nil {
		fmt.Println("PowBroadcast 服务调用失败")
	}
	fmt.Println("PowBroadcast返回",re.PowRspMsg)

	//通过句柄调用函数
	re1 ,err :=cli.PosBroadcast(context.Background(),&pd.PosReqMsg{PosReqMsg:"Rona"})
	if err !=nil{
		fmt.Println("PosBroadcast 调用失败")
	}
	fmt.Println("调用PosBroadcast的返回",re1.PosRspMsg)
}
//节点注册到主节点
func (c *Node)Register(cli pd.MainNodeServiceClient)  {
	re, err := cli.RegisterNode(context.Background(),&pd.RegisterReq{NodePort:c.Port,Token:c.Token})
	if err != nil{
		fmt.Println("节点注册失败")
	}
	fmt.Println(c.Port+"注册成功",re.RegisterRsp)
}
//挖矿函数
//DT = "00"
func (this *Node)Mining(DT string) string {
	//1.创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	for {
		//2.输入数据
		noce := rand.Int63n(100000000)   //以一个随机数为输入数据
		//fmt.Println(c.Port,"正在挖矿...")
		hash.Write([]byte(string(noce)))
		//3.计算哈希值
		s := hash.Sum(nil)
		//4.将字符串编码为16进制格式,返回字符串
		hashcode := hex.EncodeToString(s)
		//对第一次求得的hash进行二次hash计算
		hash.Write([]byte(hashcode))
		s1 := hash.Sum(nil)
		hashcode1 := hex.EncodeToString(s1)
		//打印挖矿过程
		fmt.Println("Hash256:"+hashcode1)

		if strings.HasPrefix(hashcode1,DT){
			fmt.Println("客户端"+this.Port+"挖矿成功:",hashcode1,"\n时间：",time.Now())
			tmp := this.ConnectTo("127.0.0.1:10010")

			req := pd.LedgerReq{
				Nonce:DT,
				CostTime:"99",
				Result:hashcode1,
				Ip:this.Ip + this.Port,
			}

			tmp.RecordPeer(context.Background(),&req)
			return this.Port+":"+hashcode1
		}
		time.Sleep(10*time.Millisecond)
	}
	return "default"
}