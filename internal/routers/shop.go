package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initShop(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	v1.Use(middleware.JWTAuth())
	{
		// 根据id查询店铺
		v1.GET("/shop/:id", logic.GetShopByID)
		// 全部店铺
		v1.GET("/shops", logic.GetShops)
		// 推荐店铺
		v1.GET("/shops/random/:num", logic.GetRandomShops)
		// 新增店铺
		v1.POST("/shop", logic.AddShop)
		// 注销店铺
		v1.DELETE("/shop/:id", logic.DelShop)
	}
}
