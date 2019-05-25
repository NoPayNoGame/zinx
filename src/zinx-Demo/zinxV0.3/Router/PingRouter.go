package Router

import (
	"GoServer/zinx/ziface"
	"fmt"
)

type PingRouter struct {
	ziface.IRouter
}

//	处理业务前
func (pt *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("[PreHandle] PingRouter run!")
	_, err := req.GetZConnection().GetTcpConnection().Write([]byte("[PreHandle] PingRouter run!"))
	if err != nil {
		fmt.Println(err)
	}
}

//	处理业务函数
func (pt *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("[PreHandle] Handle run!")
	_, err := req.GetZConnection().GetTcpConnection().Write([]byte("[PostPreHandle] Handle run!"))
	if err != nil {
		fmt.Println(err)
	}

}

//	处理业务后
func (pt *PingRouter) PostPreHandle(req ziface.IRequest) {
	fmt.Println("[PreHandle] PostPreHandle run!")
	_, err := req.GetZConnection().GetTcpConnection().Write([]byte("[PostPreHandle] PingRouter run!"))
	if err != nil {
		fmt.Println(err)
	}
}
