package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/routers/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.Default()
	//	中间件
	router.Use(middleware.Headers)

	router.GET("/", logic.Index)

	initUser(router)

	api := router.Group("/api")
	initLogin(api)

	initOrder(router)

	err := router.Run(":7892")
	if err != nil {
		panic(err)
	}
}
