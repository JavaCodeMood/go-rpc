package main

//https://www.jianshu.com/p/0b1f56f14c4e

import (
	"fmt"
	"os"
	"net/rpc/jsonrpc"
	"log"
)

// 类可以不一样 但是 Who 和DoWhat 要必须一样  要不然接收到不到值，
type Send struct {
	Who, DoWhat string
}

func main() {

	if len(os.Args)==4{
		fmt.Println("长度必须等于4,一个ip的地址ip=",os.Args[1],"加上后面的被除数os.Args[2]=",os.Args[2],"和除数os.Args[3]=",os.Args[3])
		//os.Exit(1)
	}

	service:=os.Args[1]
	client,err:=jsonrpc.Dial("tcp",service)
	if err != nil {
		log.Fatal("Dial 发生了错误了err=",err)
	}
	send:=Send{os.Args[2],os.Args[3]}
	var  resive  string
	err1:=client.Call("DemoM.DoWork",send,&resive)
	if err1!=nil {
		fmt.Println("shiming call error    ")
		fmt.Println("Call 的时候发生了错误了 err=",err1)
	}
	fmt.Println("收到信息了",resive)

}

//进入当前目录下，开启两个命令行窗口，一个运行Server.go, 一个运行Client.go
//执行：go run Client.go 127.0.0.1:8080 xiaoming music

