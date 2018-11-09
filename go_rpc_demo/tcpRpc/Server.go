package main

import (
	"go_rpc_demo/tcpRpc/common"
	"net/rpc"
	"fmt"
	"net"
)

func main(){
	//实例化服务对象
	var ms = new(common.MathService)
	//注册这个服务
	rpc.Register(ms)
	fmt.Println("服务启动。。。。")
	//定义TCP的服务承载地址
	var address,_ = net.ResolveTCPAddr("tcp","127.0.0.1:1234")
	//监听TCP连接
	listener,err := net.ListenTCP("tcp",address)
	if err != nil {
		fmt.Println("服务启动失败！", err)
	}
	//死循环
	for{
		//如果接收到连接
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("接收到一个调用请求。。。")
		//绑定rpc到tcp连接上
		rpc.ServeConn(conn)
	}
}