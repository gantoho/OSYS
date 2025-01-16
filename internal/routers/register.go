package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gin-gonic/gin"
)

func initRegister(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.POST("/register", logic.Register)

	}
}
