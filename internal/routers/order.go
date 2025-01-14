package routers

import "github.com/gin-gonic/gin"

func initOrder(r *gin.Engine) {
	r.GET("/order")
}
