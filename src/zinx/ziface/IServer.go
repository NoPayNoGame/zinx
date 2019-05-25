package ziface

/*
Server模块的抽象方层	(方法)
*/

type IServer interface {
	//	启动
	Start()

	//	停止
	Stop()

	//	开启服务
	Serve()

	//	添加路由
	AddRouter(router IRouter)
}
