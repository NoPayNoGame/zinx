package ziface

/*
	抽象IRequest一次性请求的数据封装
*/

type IRequest interface {
	//	得到当前请求的连接
	GetConnection() IConnection

	//	得到连接数据
	GetData() []byte

	//	得到数据长度
	GetDataLen() int
}
