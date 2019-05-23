package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	IpVersion string
	IP        string
	Port      int
	Name      string
}

//type HandleFunc func(*net.TCPConn,[]byte,int) error
func CallBackBusi(request ziface.IRequest) error {
	//	回显业务
	fmt.Println("[conn handle]CallBack..")
	conn := request.GetConnection().GetTCPConnection()
	data := request.GetData()
	cnt := request.GetDataLen()

	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write Call Back error:", err)
		return err
	}
	return nil

}

func NewServer(Name string) ziface.IServer {
	s := &Server{
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7000,
		Name:      Name,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("【start】服务器监听IP:%s,Port:%d,运行..\n", s.IP, s.Port)

	//	1 创建套接字:得到一个TCP的addr
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}

	//	2 监听服务器地址
	listenner, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("listen:", s.IpVersion, "error:", err)
		return
	}

	//	生成id累加器
	var cid uint32
	cid = 0

	//	3 阻塞等待客户端发送请求
	go func() {
		for {
			//	阻塞等待客户端请求
			conn, err := listenner.AcceptTCP() //只针对TCP协议,所以用AcceptTCP
			if err != nil {
				fmt.Println("Accept error:", err)
				continue
			}

			//	创建一个Connction对象
			//	func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) ziface.IConnection {
			//	func CallBackBusi(conn *net.TCPConn, data []byte, cnt int) error {

			//	将原生的conn 和 CallBack 绑定
			dealConn := NewConnection(conn, cid, CallBackBusi)

			cid++

			go dealConn.Start()

			//	conn和对端连接成功
			//go func() {
			//	//	4 客户端有数据请求,处理客户端业务(读,写)
			//	//	循环读写
			//	for {
			//
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("recv buf err", err) //EOF
			//			//	连接断开,不再循环
			//			break
			//		}
			//		fmt.Printf("收到客户端数据%s,长度为:%d\n", buf, cnt)
			//
			//		//	回显功能	(业务)
			//
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("输出buf错误:", err)
			//			continue
			//		}
			//	}
			//}()
		}
	}()
}
func (s *Server) Stop() {
	// TODO 将一些服务器资源进行回收

}
func (s *Server) Server() {
	//	启动server的监听功能
	s.Start()

	// TODO	做一些其他的扩展
	//	阻塞
	//	告诉CPU不需要处理,节省cpu资源
	select {}
}
