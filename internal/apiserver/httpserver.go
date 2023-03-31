package apiserver

import (
	"fmt"
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
		server.WithHandleMethodNotAllowed(true),
		server.WithHostPorts(fmt.Sprintf(":%s", config.Config.Web.Port)),
		//server.WithTransport(standard.NewTransporter),
	)
	//是否跨域
	if config.Config.Web.Cors {
		h.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	}
	//url := swagger.URL("http://localhost:9681/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	{
		v1.InitApi(h) //v1版本相关接口
	}

	//启动gin路由服务
	log.Println("端口号：", config.Config.Web.Port)
	h.Spin()
}
