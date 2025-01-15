package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gin-gonic/gin"
)

func initUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		//获取用户信息
		v1.GET("/user/:id")
		//新增用户
		v1.POST("/user", logic.Register)
		//编辑用户
		v1.PUT("/user/:id", logic.Edit)
		//删除用户
		v1.DELETE("/user/:id")
	}
}
