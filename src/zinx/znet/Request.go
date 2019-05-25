package znet

import "GoServer/zinx/ziface"

type Request struct {
	zConn   ziface.IConnection
	data    []byte
	dataLen int
}

func NewRequest(conn ziface.IConnection, data []byte, dataLen int) ziface.IRequest {
	req := &Request{
		zConn:   conn,
		data:    data,
		dataLen: dataLen,
	}
	return req
}

//	获取数据长度
func (r *Request) GetDataLen() int {
	return r.dataLen
}

//	获取数据内容
func (r *Request) GetData() []byte {
	return r.data
}

//	获取封装conn
func (r *Request) GetZConnection() ziface.IConnection {
	return r.zConn
}
