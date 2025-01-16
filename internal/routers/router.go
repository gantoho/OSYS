package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.Default()
	//	中间件
	router.Use(middleware.Headers)

	api := router.Group("/api")
	{
		api.GET("/", logic.Index)
	}
	initUser(api)
	initLogin(api)
	initOrder(api)
	initToken(api)

	err := router.Run(":7892")
	if err != nil {
		panic(err)
	}
}
