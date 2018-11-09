package main

import (
	"net/rpc"
	"fmt"
	"go_rpc_demo/tcpRpc/common"
)

func main() {
	var client,err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接服务器失败....", err)
	}
	var args = common.Args{40, 9}
	var result = common.Result{}
	fmt.Println("开始调用....")
	fmt.Println("运算参数：", args.A, args.B)

	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Println("调用失败...",err)
	}
	fmt.Println("加法结果：", result.Value)

	err = client.Call("MathService.Divide", args, &result)
	if err != nil {
		fmt.Println("调用失败...", err)
	}
	fmt.Println("除法结果：", result.Value)

	err = client.Call("MathService.Muil", args, &result)
	if err != nil {
		fmt.Println("调用失败...", err)
	}
	fmt.Println("乘法结果：", result.Value)

	err = client.Call("MathService.Subt", args, &result)
	if err != nil {
		fmt.Println("调用失败...", err)
	}
	fmt.Println("减法结果：", result.Value)
}


//需要开启两个命令行窗口
//go run Server.go
//go run Client.go
