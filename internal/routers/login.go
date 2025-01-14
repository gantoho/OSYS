package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gin-gonic/gin"
)

func initLogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/login", logic.Login)
	}
}
