package ziface

import "net"

type IConnection interface {
	//	启动连接
	Start()

	//	停止连接
	Stop()

	//	获取原生socket
	GetTcpConnection() *net.TCPConn

	//	获取连接ID
	GetTcpConnID() uint32

	//	查看对端客户端的IP和端口
	GetRemoteAddr() net.Addr

	//	读取数据.将数据发送给对端
	Send(data []byte, readLen int) (writeLen int, err error) //	TODO	暂时没用上
}
