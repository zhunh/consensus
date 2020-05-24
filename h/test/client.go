package main

import (
	"fmt"
	"google.golang.org/grpc"
	pd "consensus/grpc"
	"context"
)

func main() {
	//客户端连接服务器
	conn ,err :=grpc.Dial("127.0.0.1:10088",grpc.WithInsecure())
	if err!=nil {
		fmt.Println("网络异常",err)
	}
	//网络延迟关闭
	defer  conn.Close()


	//获得grpc句柄
	c:=pd.NewNodeServerClient(conn)

	//通过句柄调用函数
	re ,err :=c.PowBroadcast(context.Background(),&pd.PowReqMsg{PowReqMsg:"Kristin"})
	if err!=nil {
		fmt.Println("sayhello 服务调用失败")
	}
	fmt.Println("调用sayhello的返回",re.PowRspMsg)



	re1 ,err :=c.PosBroadcast(context.Background(),&pd.PosReqMsg{PosReqMsg:"Rona"})
	if err !=nil{
		fmt.Println("say name 调用失败")
	}
	fmt.Println("调用Sayname的返回",re1.PosRspMsg)

}
