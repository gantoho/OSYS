package routers

import "github.com/gin-gonic/gin"

func initOrder(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.GET("/order")
	}
}
