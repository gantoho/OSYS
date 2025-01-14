package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/routers/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	//	中间件
	router.Use(middleware.Headers)

	router.GET("/", logic.Index)

	user := router.Group("/user")
	{
		user.POST("/register", logic.Register)
		user.POST("/login", logic.Login)
	}

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
