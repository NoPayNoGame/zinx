package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	network := "tcp4"
	address := "127.0.0.1:7000"

	fmt.Println("[Client] 与服务器", address, "建立", network, "连接!")

	//	与服务器建立连接
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("[Client] Dial", network, address, "error")
		return
	}

	fmt.Println("[Client] 连接服务器成功!")

	data := "Hello Zinx!"

	for {
		time.Sleep(1 * time.Second)
		//	向服务器发送数据
		cnt, err := conn.Write([]byte(data))
		if err != nil {
			if cnt == 0 {
				fmt.Println("[Client]", conn.RemoteAddr(), "与服务器断开!")
				break
			} else {
				fmt.Println("[Client] write error:", err, cnt)
				continue
			}
		}

		//	从服务器读取数据
		buf := make([]byte, 512)
		cnt, err = conn.Read(buf)
		if err != nil {
			if cnt == 0 || err == io.EOF {
				fmt.Println("[Start]", conn.RemoteAddr(), "与服务器断开!")
				break

			} else {
				fmt.Println("[Client] Read error:", err)
				break
			}

		}
		fmt.Printf("[Client] 从服务器读到:%s,长度:%d\n", string(buf[:cnt]), cnt)

	}
}
