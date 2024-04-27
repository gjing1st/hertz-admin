package config

import "sync"

var once sync.Once

func Init() {
	once.Do(initConf)
}

func initConf() {
	InitConfig()
	InitLogger(Config.Log.Level, Config.Log.Output, Config.Log.Dir, Config.Log.Caller)
}
