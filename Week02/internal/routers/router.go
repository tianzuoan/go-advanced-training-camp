package routers

import (
	"github.com/gin-gonic/gin"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strconv"
	"week02/global"
	v1 "week02/internal/routers/api/v1"
	"week02/internal/service"
)

func NewHttpServer() error {
	gin.SetMode(gin.DebugMode)
	addr := "127.0.0.1:" + strconv.Itoa(global.ServerSetting.HttpPort)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", service.PingService)
	r.GET("/api/v1/get-user-list", v1.GetUserList)
	r.GET("/api/v1/get-user-detail", v1.GetUserDetail)
	return r.Run(addr) // listen and serve on 0.0.0.0:8088 (for windows "localhost:8088")
}
