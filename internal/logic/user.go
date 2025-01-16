package logic

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gantoho/osys/internal/models"
	"github.com/gantoho/osys/internal/tools"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type registerUser struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Password2 string `json:"password2" form:"password2" binding:"required"`
	Email     string `json:"email" form:"email"`
}

// 用户注册
func Register(c *gin.Context) {
	var rUser registerUser
	err := c.ShouldBind(&rUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "server error: " + err.Error(),
			Data:    nil,
		})
		return
	}

	fmt.Printf("rUser: %#v\n", rUser)

	if rUser.Username == "" || rUser.Password == "" || rUser.Password2 == "" {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "param error: username and password is empty",
			Data:    nil,
		})
		return
	}

	// 用户名只能包含字母、数字、下划线、横线, 长度大于3小于18
	regexUser := regexp.MustCompile("^[a-zA-Z0-9_-]{3,18}$")
	if !regexUser.MatchString(rUser.Username) {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "param error: username is invalid",
			Data:    nil,
		})
		return
	}

	// 密码必须由包含大写字母，小写字母，数字，下划线_，短横线
	//passwordRegex := regexp.MustCompile("^.(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[-_]).{6,20}$")
	if !tools.CheckPassword(rUser.Password) {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "param error: password is invalid",
			Data:    nil,
		})
		return
	}

	if rUser.Password != rUser.Password2 {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "param error: password and password2 different",
			Data:    nil,
		})
		return
	}

	var user models.User
	_ = models.DB.Where("username = ?", rUser.Username).First(&user).Error
	if user.ID > 0 {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "username is exist",
			Data:    nil,
		})
		return
	}

	err = models.DB.Create(&models.User{
		Username: rUser.Username,
		Password: tools.Encrypt(rUser.Password),
		Email:    rUser.Email,
		Model: models.Model{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "gorm err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    gin.H{"username": rUser.Username},
	})
}

type loginUser struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var lUser loginUser
	err := c.ShouldBind(&lUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "server error: " + err.Error(),
		})
		return
	}

	var user models.User
	err = models.DB.Where("username = ?", lUser.Username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "gorm err: " + err.Error(),
		})
		return
	}

	if user.Password != tools.Encrypt(lUser.Password) {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "param error: password err",
		})
		return
	}

	// 登录成功，创建token
	newJwt := middleware.NewJWT()
	token, _ := newJwt.CreateToken(middleware.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    user.Username,
			NotBefore: time.Now().Unix(),
		},
	})

	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    gin.H{"token": token, "user": user},
	})
}

func Edit(c *gin.Context) {
	//var edit_user models.User
	//err := c.ShouldBind(&edit_user)
	//if err != nil {
	//	panic("logic user edit ShouldBind err")
	//}
	//
	//userID, _ := strconv.Atoi(c.Param("id"))
	//edit_user.ID = uint(userID)
	//
	//var user models.User
	//
	//models.DB.Find(&user, "id = ?", userID).Updates(
	//	map[string]interface{}{
	//		"username": edit_user.Username,
	//		"password": edit_user.Password,
	//		"phone":    edit_user.Phone,
	//	},
	//)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "edit success",
	//})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	err := models.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "gorm err: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data:    gin.H{"user": user},
	})
}
