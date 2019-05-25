package znet

import (
	"GoServer/zinx/config"
	"GoServer/zinx/ziface"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	//	原生socket
	Conn *net.TCPConn

	//	connID
	ConnID uint32

	//	记录连接是否关闭
	IsClosed bool

	//	当前连接绑定的函数
	//handleAPI HandleFunc

	Router ziface.IRouter
}

//Send(data []byte, readLen int) (writeLen int, err error)
type HandleFunc func(req ziface.IRequest) (outLen int, err error)

//		ConnID 		自增长
//		IsClosed	默认false
func NewConnection(Conn *net.TCPConn, ConnID uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		Conn:     Conn,
		ConnID:   ConnID,
		IsClosed: false,
		//handleAPI: CallBackAPI,
		Router: router,
	}

	return c
}

//	开始读数据
func (c *Connection) StartReader() {
	fmt.Println("[StartReader] goroutine is start..")
	defer c.Stop()

	for {
		buf := make([]byte, config.GlobalObject.MaxPackageSize)
		inLen, err := c.Conn.Read(buf)
		if err != nil {

			if inLen == 0 || err == io.EOF {
				fmt.Println("[StartReader] ConnId:", c.ConnID, "IP:", c.Conn.RemoteAddr(), " is down!")
				break
			} else {
				fmt.Println("[StartReader] recv buf error:", err)
				break
			}
		}
		fmt.Printf("[StartReader] 从 %s 收到: %s 长度: %d\n", c.Conn.RemoteAddr().String(), buf[:inLen-1], inLen)

		//_, _ = c.Send(buf, inLen)		//	TODO	Send方法暂时没用上
		req := NewRequest(c, buf, inLen)

		c.Router.PreHandle(req)
		c.Router.Handle(req)
		c.Router.PostPreHandle(req)

		//if _, err := c.handleAPI(req); err != nil {
		//	fmt.Println("[StartReader] ConnId:", c.ConnID, "Handle error:", err)
		//	break
		//}
	}

}

//	启动连接
func (c *Connection) Start() {
	go c.StartReader()

	//	TODO	其他操作
}

//	停止连接
func (c *Connection) Stop() {
	fmt.Println("[Stop] ConnId:", c.ConnID)
	if !c.IsClosed {
		c.IsClosed = true
	}
	_ = c.Conn.Close()
}

//	获取原生socket
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

//	获取连接ID
func (c *Connection) GetTcpConnID() uint32 {
	return c.ConnID
}

//	查看对端客户端的IP和端口
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//	读取数据.将数据发送给对端		//	TODO	暂时没用上
func (c *Connection) Send(data []byte, inLen int) (outLen int, err error) {
	if outLen, err := c.Conn.Write(data[:inLen]); err != nil {
		return outLen, err
	}
	return outLen, nil
}
