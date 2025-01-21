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
	initRegister(api)
	initLogin(api)
	initUser(api)
	initOrder(api)
	initToken(api)
	initProduct(api)
	initShop(api)
	initUpload(api)

	err := router.Run(":7892")
	if err != nil {
		panic(err)
	}
}
