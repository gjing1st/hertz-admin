package config

func Init() {
	InitConfig()
	InitLogger(Config.Log.Level, Config.Log.Output, Config.Log.Dir, Config.Log.Caller)
}
