package main

import (
	"net/rpc"
	"fmt"
	"net/http"
	"go_rpc_demo/httpRpc/common"
)

func main(){
	//新建一个服务
	var ms = new(common.MathService)
	//注册服务
	rpc.Register(ms)
	//将PRC绑定到Http协议上
	rpc.HandleHTTP()
	fmt.Println("启动服务。。。。")
	//监听服务端口
	err := http.ListenAndServe(":1234", nil)
	if err != nil{
		fmt.Println(err.Error)
	}
	fmt.Println("服务已经停止")

}
