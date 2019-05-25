package znet

import "GoServer/zinx/ziface"

type Router struct {
}

//	处理业务前
func (rt *Router) PreHandle(req ziface.IRequest) {

}

//	处理业务函数
func (rt *Router) Handle(req ziface.IRequest) {

}

//	处理业务后
func (rt *Router) PostPreHandle(req ziface.IRequest) {

}
