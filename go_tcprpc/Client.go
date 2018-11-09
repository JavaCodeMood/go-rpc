package main

//https://www.jianshu.com/p/0b1f56f14c4e

import (
	"fmt"
	"os"
	"net/rpc"
	"log"
	"strconv"
)

func main() {

	if len(os.Args)==4{
		fmt.Println("长度必须等于4,一个ip的地址ip=",os.Args[1],"加上后面的被除数os.Args[2]=",os.Args[2],"和除数os.Args[3]=",os.Args[3])
		//os.Exit(1)
	}
	// 获取 ip 地址
	service:= os.Args[1]
	//连接 拨号连接到指定的网络地址的RPC服务器。
	client,err:=rpc.Dial("tcp",service)
	if err!=nil {
		log.Fatal("连接过程中发生错误退出",err)
	}
	num1:=os.Args[2]
	i1,error1:=strconv.Atoi(num1)
	if error1!=nil {
		fmt.Println("error ：",error1)
		os.Exit(1)
	}
	num2:=os.Args[3]
	i2,error2:=strconv.Atoi(num2)
	if error2!=nil {
		fmt.Println("error ：",error2)
		os.Exit(1)
	}
	aa:=AAA{i1,i2}
	var reply  int
	err1:=client.Call("Num.M",aa,&reply)

	if err1 != nil{
		log.Fatal("Call的时候发生了错误",err1)
	}
	fmt.Printf("Num : %d*%d=%d\n",aa.A,aa.B,reply)

	var bb BDemo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err= client.Call("Num.F",aa,&bb)
	if err!=nil {
		log.Fatal("err=====",err)
	}
	fmt.Printf("Num: %d/%d=%d 余数 %d\n",aa.A,aa.B,bb.DD,bb.CC)

}


// 定义两个类，那边需要操作的类
type AAA struct {
	A,B int
}
//记住这里不能够大写 两个连着一起大写 有点意思
//reading body gob: type mismatch: no fields matched compiling decoder for  DDDD
//todo 为啥 第二个参数  只要是两个连在一起的DDDD   就会报错   reading body gob: type mismatch: no fields matched compiling decoder for
type BDemo struct {
	DD, CC int
}


//进入当前目录下面,开启两个命令窗口，一个运行Server.go,一个运行Client.go
//执行：go run Client.go 127.0.0.1:1234 40 4