package Peers

import (
	"fmt"
	"io"
	"net/http"
)
//存放节点的名字节点的地址
var PeersTable=make(map[string]string)


//代表小国家
type Peers struct {
	//节点名称
	Id string
	//路径
	Path string
	//http响应
	Writer http.ResponseWriter
}
//先初始化部分
func InitPeers(id string,path string) *Peers{
	var newpeer Peers
	newpeer.Id=id
	newpeer.Path=path
	return &newpeer
}

//当http服务器接收到网络请求，带/req,则用此方法
//http的回调函数http://localhost:1111/req?warTime=hello
func(peer *Peers)Request(writer http.ResponseWriter,request *http.Request){
	//接受请求数据
	request.ParseForm()
	if(len(request.Form["warTime"])>0){
		peer.Writer=writer
		fmt.Println("aaa",writer)
		fmt.Println("主节点接受到的参数信息为",request.Form["warTime"][0])
		peer.BoradCast(request.Form["warTime"][0],"/prePrepare")
	}


}
//节点开始广播
func(peer *Peers)BoradCast(msg string,path string){
	fmt.Println(path)
	for peerId,url :=range PeersTable{
		//无需向自己进行分发
		if peerId==peer.Id{
			continue
		}
		//请求广播
		http.Get("http://"+url+path+"?warTime="+msg+"&peerId="+peer.Id)
	}
}
//第二阶段
func(peer *Peers)PrePrepare(writer http.ResponseWriter,request *http.Request){
	//接受请求数据
	request.ParseForm()
	if(len(request.Form["warTime"])>0){
		fmt.Println("子节点接受到主节点的广播参数信息为",request.Form["warTime"][0])
		peer.BoradCast(request.Form["warTime"][0],"/prepare")
	}


}
//存放校验成功的节点
var AuthenticationPeerMap=make(map[string]string)
var AuthenticationSuceess=false
//第三阶段
func(peer *Peers)Prepare(writer http.ResponseWriter,request *http.Request){
	//接受请求数据
	request.ParseForm()
	if(len(request.Form["warTime"])>0){
		fmt.Println("子节点接受到其余自己节点的广播参数信息为",request.Form["warTime"][0])
		//进行拜占庭校验
		peer.Authentication(request)

	}


}
//校验函数
func(peer *Peers)Authentication(request *http.Request){
	if !AuthenticationSuceess{
		if len(request.Form["peerId"][0])>0{
			AuthenticationPeerMap[request.Form["peerId"][0]]="ok"
			if len(AuthenticationPeerMap)>(len(PeersTable)-1)/3{
				AuthenticationSuceess=true
				peer.BoradCast(request.Form["warTime"][0],"/commit")
			}
		}
	}



}

func(peer *Peers)Commit(writer http.ResponseWriter,request *http.Request){
	//接受请求数据

	if writer!=nil{
		fmt.Println("拜占庭校验成功")
		io.WriteString(peer.Writer,"ok")
	}


}

