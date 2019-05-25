package ziface

type IRouter interface {
	//	处理前
	PreHandle(req IRequest)

	//	处理函数
	Handle(req IRequest)

	//	处理请求后
	PostPreHandle(req IRequest)
}
