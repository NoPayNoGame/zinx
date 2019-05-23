package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

//具体的TCP链接模块
type Connection struct {
	//当前链接的原生套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前的链接状态
	isClosed bool

	//当前链接所绑定的业务处理方法
	handleAPI ziface.HandleFunc
}

/*
初始化链接方法HandleFunc
*/
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) ziface.IConnection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
	}

	return c
}

//	开始读取,并调用HandleAPI
func (c *Connection) StartReader() {
	fmt.Println("Start Reader is start...")
	buf := make([]byte, 512)
	for {
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err:", err)
			break
		}
		//将数据 传递给我们 定义好的Handle Callback方法
		//xxx	type HandleFunc func(*net.TCPConn,[]byte,int) error

		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req := NewRequest(c, buf, cnt)

		//	讲述传递给我们定义好的Handle CallBack方法s
		//c.handleAPI(req)

		err = c.handleAPI(req)
		if err != nil {
			fmt.Println("ConnID", c.ConnID, "Handle is error", err)
			break
		}
	}
}

//启动链接
func (c *Connection) Start() {
	fmt.Println("conn start(),conn ID =", c.ConnID)
	go c.StartReader()
}

//停止链接
func (c *Connection) Stop() {
	fmt.Println("conn stop(),conn ID =", c.ConnID)
	if c.isClosed == true {
		return
	}

	c.isClosed = true
	//	关闭原生套接子
	_ = c.Conn.Close()
}

//获取链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取conn的原生socket套接字
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取远程客户端的ip地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给对方客户端
func (c *Connection) Send(data []byte, cnt int) error {
	if _, err := c.Conn.Write(data[:cnt]); err != nil {
		fmt.Println("Send buf error:")
		return err
	}
	return nil
}
