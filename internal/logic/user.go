package logic

import (
	"fmt"
	"github.com/gantoho/osys/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		panic("logic user register ShouldBind err")
	}

	err = model.DB.Create(&user).Error
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

	var user model.User
	err = model.DB.Where("username = ?", login_user.Username).First(&user).Error
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
