package main

import "GoServer/zinx/znet"

func main() {
	//	创建一个serve对象
	serve := znet.NewServer("蓝月传奇", "tcp4", "0.0.0.0", 7000)

	//	调用启动方法
	serve.Serve()

}
