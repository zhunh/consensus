package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//简单模拟PBFT系统节点信息同步校验

//声明nodeInfo节点，代表各个小国家
type nodeInfo struct {
	//节点名称
	id string
	//路径
	path string
	//http响应
	writer http.ResponseWriter
}

//存系欸但名字和节点地址
var nodeTable = make(map[string]string)

func main() {
	//接收终端传来的参数
	userId := os.Args[1]
	fmt.Println(userId)

	//存储四个国家的地址
	nodeTable = map[string]string{
		"N0": "localhost:1111",
		"N1": "localhost:1112",
		"N2": "localhost:1113",
		"N3": "localhost:1114",
	}

	//创建国家
	node := nodeInfo{id: userId, path: nodeTable[userId]}

	//http协议的回调函数
	//http://localhost:1111/req?warTime=hello
	http.HandleFunc("/req", node.request)
	http.HandleFunc("/prePrepare", node.prePrepare)
	http.HandleFunc("/prepare", node.prepare)
	http.HandleFunc("/commit", node.commit)

	//启动监听
	if err := http.ListenAndServe(node.path, nil); err != nil {
		fmt.Println(err)
	}
}

//当http服务器接收到网络请求，带/req，则回调此方法
func (node *nodeInfo) request(writer http.ResponseWriter, request *http.Request) {
	//接收请求数据
	request.ParseForm()
	if (len(request.Form["warTime"]) > 0) {
		//证明接收到了客户端发来的数据
		node.writer = writer
		fmt.Println("主节点接收到的参数信息为", request.Form["warTime"][0])
		//激活主节点后，将数据分发给其他节点，广播
		node.broadcast(request.Form["warTime"][0], "/prePrepare")
	}
}

//节点发送广播的方法
func (node *nodeInfo) broadcast(msg string, path string) {
	fmt.Println("广播", path)
	//遍历所有节点进行广播
	for nodeId, url := range nodeTable {
		//判断是否是自己，若为自己，则跳出档次循环
		if nodeId == node.id {
			continue
		}
		//要进行分发的节点
		//像req那样，通过http请求广播
		http.Get("http://" + url + path + "?warTime=" + msg + "&nodeId=" + node.id)
	}
}

func (node *nodeInfo) prePrepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("接收到的广播", request.Form["warTime"][0])
	//若数据有值，则进行广播
	if len(request.Form["warTime"]) > 0 {
		//有值
		node.broadcast(request.Form["warTime"][0], "/prepare")
	}
}

func (node *nodeInfo) prepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("接收到子节点的广播", request.Form["warTime"][0])
	//数据大于0，认为接收到
	if len(request.Form["warTime"]) > 0 {
		//进行拜占庭校验
		node.authentication(request)
	}
}

//用于记录正常响应的好的节点
var authenticationNodeMap = make(map[string]string)
//定义标签
var authenticationSuccess = false
//拜占庭校验
func (node *nodeInfo) authentication(request *http.Request) {
	//第一次进去
	if !authenticationSuccess {
		//运行的正常节点的判断
		if len(request.Form["nodeId"][0]) > 0 {
			//证明节点是OK的
			authenticationNodeMap[request.Form["nodeId"][0]] = "OK"
			//如果有两个国家节点正确返回了结果，可以发送
			if len(authenticationNodeMap) > len(nodeTable)/3 {
				//(n-1)/3的容错性
				//进入commit阶段
				authenticationSuccess = true
				node.broadcast(request.Form["warTime"][0], "/commit")
			}
		}
	}
}

func (node *nodeInfo) commit(writer http.ResponseWriter, request *http.Request) {
	if writer != nil {
		fmt.Println("拜占庭校验成功")
		io.WriteString(node.writer, "ok")
	}
}
