package logic

import (
	"fmt"
	"github.com/gantoho/osys/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		panic("logic user register ShouldBind err")
	}

	fmt.Printf("%+v", user)

	err = models.DB.Create(&user).Error
	if err != nil {
		panic("logic user register gorm create err")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "register success",
	})
}

type loginUser struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var login_user loginUser
	err := c.ShouldBind(&login_user)
	if err != nil {
		panic("logic user login ShouldBind err")
	}
	fmt.Printf("1 %+v", login_user)

	var user models.User
	err = models.DB.Where("username = ?", login_user.Username).First(&user).Error
	fmt.Printf("2 %+v", user)
	if err != nil {
		panic("username undefined")
	}
	if user.Password != login_user.Password {
		panic("password error")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "login success",
	})
}
