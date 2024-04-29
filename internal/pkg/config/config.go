package config

import (
	"github.com/jinzhu/configor"
)

// Config 存储全局参数，供其他模块使用
var Config = struct {
	Base        Base
	Log         Log
	Database    Database
	Slave       Slave
	VersionInfo VersionInfo
	CrontabTime int `default:"60"`

	//网卡相关配置
	Adapter struct {
		AdminPath  string `default:"/etc/sysconfig/network-scripts/ifcfg-eno1"`
		CipherPath string `default:"/etc/sysconfig/network-scripts/ifcfg-eno1"`
	}
	//升级包存放的目录
	UploadPath string `default:"/opt/tnaengine/update/"`
	AssistAddr string `default:"http://172.17.0.1:18998"`
}{}

type Base struct {
	DBType          string `default:"mysql"`
	CacheType       string `default:"gcache"`
	Port            string `default:"9680"`
	EnableIntegrity bool   `default:"true"`
	PwdMaxErrNum    int    `default:"5"`
}
type Log struct {
	Output string `default:"std"`  //日志输出，标准输出或文件
	Level  string `default:"info"` //日志等级
	Caller bool   `default:"true"` //是否打印调用者信息
	Dir    string `default:"."`    //存放目录
}

type Database struct {
	Host     string `default:"localhost"`
	UserName string `default:"root"`
	Password string `default:"123456"`
	DBName   string `default:"ha"`
	Port     string `default:"3306"`
	MinConns int    `default:"90"`  //连接池最大空闲连接数量 不要太小
	MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
	SSLMode  string `default:"disable"`
}

type Slave struct {
	Host     string `default:"sample-follower.third"`
	UserName string `default:"super_user"`
	Password string `default:"CSMP@1qaz2wsx"`
	DBName   string `default:"alert"`
	Port     string `default:"3306"`
	MinConns int    `default:"90"`  //连接池最小连接数量 不要太小
	MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
}
type VersionInfo struct {
	Manufacturer string `default:"xxxx"`
	Serial       string `default:"35D485H3B7Z89N"`
	DeviceModel  string `default:"serial1212345678"`
	Version      string `default:"1.0.0"`
	//Algorithm    string `default:"SM2、SM3、SM4"`
}

// InitConfig 读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./configs/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
}
