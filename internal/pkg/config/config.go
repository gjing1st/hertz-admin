package config

import (
	"github.com/jinzhu/configor"
)

// Config 存储全局参数，供其他模块使用
var Config = struct {
	Log struct {
		Output string `default:"std"`  //日志输出，标准输出或文件
		Level  string `default:"info"` //日志等级
		Caller bool   `default:"true"` //是否打印调用者信息
		Dir    string `default:"."`    //存放目录
	}
	Web struct {
		Port string `default:"9680"`
		Cors bool   `default:"true"`
	}
	Mysql struct {
		Host     string `default:"114.115.134.131"`
		UserName string `default:"root"`
		Password string `default:"ZPFIZgvCev"`
		DBName   string `default:"alert"`
		Port     string `default:"30324"`
		MinConns int    `default:"90"`  //连接池最大空闲连接数量 不要太小
		MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
	}
	CrontabTime int `default:"60"`
	VersionInfo struct {
		Manufacturer string `default:"xxxx"`
		Serial       string `default:"35D485H3B7Z89N"`
		DeviceModel  string `default:"serial1212345678"`
		Version      string `default:"1.0.0"`
		//Algorithm    string `default:"SM2、SM3、SM4"`
	}

	//网卡相关配置
	Adapter struct {
		AdminPath  string `default:"/etc/sysconfig/network-scripts/ifcfg-eno1"`
		CipherPath string `default:"/etc/sysconfig/network-scripts/ifcfg-eno1"`
	}
	//升级包存放的目录
	UploadPath string `default:"/opt/tnaengine/update/"`
	AssistAddr string `default:"http://172.17.0.1:18998"`
}{}

// InitConfig 读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./config/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
}
