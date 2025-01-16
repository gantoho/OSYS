package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initUser(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	v1.Use(middleware.JWTAuth())
	{
		//获取用户信息
		v1.GET("/user/:id", logic.GetUserByID)
		//编辑用户
		v1.PUT("/user/:id", logic.Edit)
		//删除用户
		v1.DELETE("/user/:id")
	}
}
