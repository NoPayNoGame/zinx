package znet

import (
	"GoServer/zinx/config"
	"GoServer/zinx/ziface"
	"fmt"
	"net"
)

/*
	Server模块的实现层(属性,方法实现,初始化New方法-->返回IServer)
*/
type Server struct {
	//	服务器名称
	Name string
	//	IP协议
	IpVersion string
	//	IP
	IP string
	//	端口
	Port uint32
	//	路由
	Router ziface.IRouter
}

func NewServer(ipVersion string) ziface.IServer {
	s := &Server{
		Name:      config.GlobalObject.Name,
		IpVersion: ipVersion,
		IP:        config.GlobalObject.Host,
		Port:      config.GlobalObject.Port,
		Router:    nil,
	}
	return s
}

//	定义一个	具体的回显业务	type HandleFunc func(conn *net.TCPConn, data []byte, inLen int) (outLen int, err error)

//	fixme	不要了
//func CallBackBusiness(req ziface.IRequest) (outLen int, err error) {
//	if outLen, err := req.GetZConnection().GetTcpConnection().Write(req.GetData()[:req.GetDataLen()]); err != nil {
//		return outLen, err
//	}
//	return outLen, nil
//}

//	启动
func (s *Server) Start() {
	fmt.Println("[Start] 开始监听...")

	//	1	创建套接字		得到一个TCP的addr
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("[Start] resolve tcp addr error:", err)
		return
	}

	//	2	监听服务器地址
	listener, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("[Start] listen tcp error:", err)
		return
	}

	//	起始id
	var cid uint32 = 0

	//	3	阻塞等待客户端发送请求
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("[Start] listener accept error:", err)
			continue
		}
		fmt.Println("[Start]", conn.RemoteAddr().String(), "连接到服务器!")

		iConn := NewConnection(conn, cid, s.Router)

		/*
			//	4	客户端有数据请求,处理客户端业务(读,写)
			//	创建goroutine	防止业务阻碍继续监听
			//go func() {
			//	//	循环读写
			//	for {
		*/

		go iConn.Start()

		/*
					//	读数据
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)

					//	读取错误
					if err != nil {
						//	断开连接
						if cnt == 0 || err == io.EOF {
							fmt.Println("[Start]", conn.RemoteAddr(), "与服务器断开!")
							break

						} else {
							fmt.Println("[Start] reader buf error:", err)
							break
						}
					}

					//	回显数据
					_, err = conn.Write(buf[:cnt])
					if err != nil {
						fmt.Println("[Start] write buf error:", err)
						break
					}
					fmt.Printf("[Start] 从 %s 收到: %s 长度: %d\n", conn.RemoteAddr().String(), buf[:cnt-1], cnt)

				}
			}()
		*/
	}

}

//	停止
func (s *Server) Stop() {
	//	TODO	回收服务器资源
}

//	开启服务
func (s *Server) Serve() {
	fmt.Println("[Serve] 服务器", s.Name, "已开启,IP:", s.IP, "端口:", s.Port)

	//	调用开始监听
	s.Start()

	//	TODO	完成其他操作

	//	阻塞
	select {}
}

//	添加路由
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}
