package version

var version string
var appName = "ha-server"

func GetAppName() string {
	return appName
}

func GetVersion() string {
	return version
}
