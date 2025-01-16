package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gin-gonic/gin"
)

func initToken(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.GET("/token", logic.GetUserIDByToken)
	}
}
