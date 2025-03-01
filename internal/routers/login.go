package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gin-gonic/gin"
)

func initLogin(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.POST("/login", logic.Login)
	}
}
