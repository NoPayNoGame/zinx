package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("客户端开始运行...")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		fmt.Println("连接服务器错误:", err)
		return
	}

	for ; ; {
		_, err := conn.Write([]byte("你好金克斯..."))
		if err != nil {
			fmt.Println("输出数据错误:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取数据错误:", err)
			return
		}

		fmt.Printf(" servar call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}
