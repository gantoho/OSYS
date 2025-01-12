package logic

import (
	"github.com/gantoho/osys/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUser(c *gin.Context) {
	f := models.NewUser()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": f,
	})
}
