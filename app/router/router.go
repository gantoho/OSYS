package router

import (
	"github.com/gantoho/osys/app/logic"
	"github.com/gantoho/osys/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	//	中间件
	router.Use(middleware.Headers)

	router.GET("/", logic.Index)

}
