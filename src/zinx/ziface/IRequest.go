package ziface

type IRequest interface {
	//	获取数据长度
	GetDataLen() int

	//	获取数据内容
	GetData() []byte

	//	获取封装conn
	GetZConnection() IConnection
}
