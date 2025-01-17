package routers

import (
	"github.com/gantoho/osys/internal/logic"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initProduct(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	v1.Use(middleware.JWTAuth())
	{
		// 根据id获取商品
		v1.GET("/product/:id", logic.GetProductByID)
		// 根据编号获取商品
		v1.GET("/product", logic.GetProductByNo)
		// 获取全部商品
		v1.GET("/products", logic.GetProduct)
		// 添加商品
		v1.POST("/product", logic.AddProduct)
		// 修改商品
		v1.PUT("/product/:id", logic.EditProduct)
		// 根据id删除商品
		v1.DELETE("/product/:id", logic.DelProduct)
	}
}
