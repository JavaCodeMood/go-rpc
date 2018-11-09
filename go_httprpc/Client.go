package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {

	fmt.Println("os*****************", os.Args, "**********************")
	fmt.Println("length=", len(os.Args))
	if len(os.Args) != 2 {
		// todo  第二个地址是  我们本地的地址
		fmt.Println("server  end")
		os.Exit(1)
	} else {
		fmt.Println("长度是多少 " + strconv.Itoa(len(os.Args)))
	}
	//获取输入的地址是获取输入得 os 数据的 第一个位置的值
	serverAddress := os.Args[1]
	fmt.Println("severAddress==", serverAddress)
	// //DelayHTTP在指定的网络地址连接到HTTP RPC服务器
	// /在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("发生错误了 DialHTTP", err)
	}
	i1, _ := strconv.Atoi(os.Args[2])
	i2, _ := strconv.Atoi(os.Args[3])
	args := ArgsTwo{i1, i2}
	var reply int
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Call Multiply  发生错误了哦   arith error:", err)
	}
	fmt.Printf("Arith 乘法: %d*%d=%d\n", args.A, args.B, reply)
	var quot QuotientTwo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith 除法取整数: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
