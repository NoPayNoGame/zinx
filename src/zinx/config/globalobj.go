package config

import (
	"encoding/json"
	"io/ioutil"
)

//	全局配置文件的类

/*
	server配置
*/
type GlobalObj struct {
	Host string //	当前监听的IP
	Port uint32 //	当前监听的Port
	Name string //	当前服务器名称

	Version        string
	MaxPackageSize uint32
}

//	定义一个全局的对外的配置的对象
var GlobalObject *GlobalObj

//	添加一个加在配置文件的方法
func (g *GlobalObj) LoadConfig() {
	data, err := ioutil.ReadFile("conf/zinx/zinx.json") //针对main主进程的相对路径
	if err != nil {
		panic(err)
	}

	//	将zinx.json的数据转换到GlobalObject
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

func init() {
	//	配置文件的读取操作
	//	先配置默认的,再读取用户设置的
	GlobalObject = &GlobalObj{
		//	默认值
		Name:           "蓝月传奇",
		Host:           "0.0.0.0",
		Port:           7000,
		Version:        "V0.4",
		MaxPackageSize: 512,
	}
	//	加在配置文件
	GlobalObject.LoadConfig()
}
