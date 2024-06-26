package apiserver

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	_ "github.com/gjing1st/hertz-admin/docs"
	v1 "github.com/gjing1st/hertz-admin/internal/apiserver/router/v1"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	"time"
)

// HttpStart
//
//	@description:	开始http服务
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2022/4/12 18:56
//	@success:
func HttpStart() {
	run()
}

// @description:	启动http服务
// @author:		GJing
// @email:			gjing1st@gmail.com
// @date:			2022/4/12 18:56
// @success:
func run() {
	h := server.Default(
		//server.WithHandleMethodNotAllowed(true),
		server.WithHostPorts(fmt.Sprintf(":%s", config.Config.Base.Port)),
		//server.WithTransport(standard.NewTransporter),
	)

	//h.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},                      // 允许所有来源
	//	AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // 允许的 HTTP 方法
	//	AllowHeaders:     []string{"*"},                      // 允许所有请求头
	//	AllowCredentials: true,                               // 允许发送身份凭证，如 Cookies
	//	MaxAge:           12 * time.Hour,                               // 预检请求的有效期
	//}))
	//注册中间件
	registerMiddleware(h)

	//url := swagger.URL("http://localhost:9681/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	{
		v1.InitApi(h) //v1版本相关接口
	}

	//启动gin路由服务
	log.Println("端口号：", config.Config.Base.Port)
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {

	//// pprof
	//if conf.GetConf().Hertz.EnablePprof {
	//	pprof.Register(h)
	//}
	//
	//// gzip
	//if conf.GetConf().Hertz.EnableGzip {
	//	h.Use(gzip.Gzip(gzip.DefaultCompression))
	//}
	//
	//// access log
	//if conf.GetConf().Hertz.EnableAccessLog {
	//	h.Use(accesslog.New())
	//}

	// recovery
	h.Use(recovery.Recovery())

	// cors
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}
