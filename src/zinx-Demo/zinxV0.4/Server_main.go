package main

import (
	"GoServer/zinx-Demo/zinxV0.3/Router"
	"GoServer/zinx/znet"
)

func main() {
	//	创建一个serve对象
	serve := znet.NewServer("tcp4")

	serve.AddRouter(&Router.PingRouter{})

	//	调用启动方法
	serve.Serve()

}
