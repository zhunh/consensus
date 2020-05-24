package node

import (
	"bufio"
	pd "consensus/grpc"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type MainNode struct {
	Port string
	RegisteredNode []Node
}
//节点注册，达到节点数开始挖矿
func (m *MainNode) RegisterNode(ctx context.Context, in *pd.RegisterReq) (*pd.RegisterRsp, error) {
	//将请求注册节点端口保存
	n := Node{Port:in.NodePort,Ledger:"default",Token:in.Token}

	m.RegisteredNode = append(m.RegisteredNode,n)
	fmt.Println("已注册的节点：",m.RegisteredNode)

	data, _ := json.Marshal(m)
	WriteIn("./consensus/nodes.txt",string(data))

	return &pd.RegisterRsp{RegisterRsp:"welcomeJoin"},nil
}
//记账权确认
func (m *MainNode) RecordPeer(ctx context.Context, in *pd.LedgerReq) (*pd.LedgerRsp, error) {
	fmt.Println("本轮获得记账权的节点为:",in.Ip,"\n目标值为：",in.Result,"\n时间为:",time.Now())

	out := pd.LedgerRsp{LedgerRsp:"sdf"}
	return &out, nil
}
//作为服务启动
func (m *MainNode)AsServer()  {
	//创建网络
	conn, err := net.Listen("tcp",":"+m.Port)
	if err != nil{
		fmt.Println("网络错误",err)
	}
	//创建grpc的服务
	srv := grpc.NewServer()

	//注册服务
	pd.RegisterMainNodeServiceServer(srv,&MainNode{})

	fmt.Println("主节点服务已启动",m.Port)

	//等待网络连接
	err = srv.Serve(conn)
	if err != nil{
		fmt.Println("网络错误",err)
	}

}
//与节点建立连接
func (m *MainNode)ConnectTo(addr string) pd.NodeServerClient {
	fmt.Println(m.Port,"开始连接",addr)
	//客户端连接服务器
	conn, err := grpc.Dial(addr,grpc.WithInsecure())
	if err != nil{
		fmt.Println("网络异常",err)
	}
	//网络延迟关闭
	//defer conn.Close()

	//获取grpc句柄 == 连接普通node
	cli := pd.NewNodeServerClient(conn)
	//通过句柄调用函数
	return cli
}
//发送挖矿指令
func (m *MainNode)MiningReq(cli pd.NodeServerClient) string {
	re, err := cli.StartMining(context.Background(),&pd.StartMiningReq{DT:"00"})
	if err != nil{
		fmt.Println("StartMining failed.",err)
	}
	//fmt.Println(re.StartMiningRsp)
	return re.StartMiningRsp
}
func (m *MainNode)StartMining(){
	b := ReadOut("./consensus/nodes.txt")
	err := json.Unmarshal(b,&m)
	if err != nil{
		fmt.Println("Unmarshal err.",err)
	}

	fmt.Println("开始挖矿..")
	fmt.Println("主节点端口：",m.Port)
	fmt.Println("子节点注册端口：",m.RegisteredNode)

	//通知各节点开始挖矿
	for k:=1;k<=2;k++{
		for i,v := range m.RegisteredNode{
			fmt.Println(i)
			fmt.Println(v.Port)
			tmp := m.ConnectTo("127.0.0.1:"+m.RegisteredNode[i].Port)
			go m.MiningReq(tmp)
		}
		fmt.Printf("开始第%d轮挖矿...\n",k)
		time.Sleep(10*time.Second)
	}
}
//pos统计权益持有情况
func (m *MainNode)Pos() Node {
	b := ReadOut("./consensus/nodes.txt")
	err := json.Unmarshal(b,&m)
	if err != nil{
		fmt.Println("Unmarshal err.",err)
	}
	var total int
	var ll []int
	//计算总的token
	for _,v := range m.RegisteredNode{
		to, _:= strconv.Atoi(v.Token)
		ll = append(ll, to)
		total += to
	}
	fmt.Println("总token为：",total)
	//使用随机值选举记账节点
	t := rand.Intn(total)
	fmt.Println("Pos随机值为：",t)
	var tmp int
	var tag int
	for m,n := range ll{
		tmp+=n
		if tmp > t{
			tag = m //记下标志位
			break
		}
	}

	return m.RegisteredNode[tag]
}
//写入文件
func WriteIn(path,s string)  {
	if _,err := os.Stat(path);os.IsNotExist(err){
		os.Remove("./consensus/nodes.txt")
		fmt.Println("已删除nodes.txt文件")
	}

	fp, err := os.Create(path)
	if err != nil{
		fmt.Println("创建文件失败")
	}
	defer fp.Close()
	fp.WriteString(s)

}
//读取文件
func ReadOut(path string) []byte {
	fp,err := os.Open(path)
	if err != nil{
		fmt.Println("打开文件错误")
		return nil
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	//读取一行内容，以`\n`为分隔符
	b,_:=r.ReadBytes('\n')

	return b
}