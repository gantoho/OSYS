package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initOrder(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	v1.Use(middleware.JWTAuth())
	{
		// 根据订单id获取订单
		v1.GET("/order/:id", logic.GetOrderByID)
		// 根据订单号获取订单
		v1.GET("/order", logic.GetOrderByNo)
		// 根据用户id获取订单
		v1.GET("/orders/:userId", logic.GetOrderByUserID)
		// 新增订单
		v1.POST("/order", logic.AddOrder)
		// 删除订单
		v1.DELETE("/order/:id", logic.DelOrder)
	}
}
