package logic

import (
	"fmt"
	"github.com/gantoho/osys/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		panic("logic user register ShouldBind err")
	}

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

func Edit(c *gin.Context) {
	var edit_user models.User
	err := c.ShouldBind(&edit_user)
	if err != nil {
		panic("logic user edit ShouldBind err")
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	edit_user.ID = uint(userID)

	var user models.User

	models.DB.Find(&user, "id = ?", userID).Updates(
		map[string]interface{}{
			"username": edit_user.Username,
			"password": edit_user.Password,
			"phone":    edit_user.Phone,
		},
	)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "edit success",
	})
}
