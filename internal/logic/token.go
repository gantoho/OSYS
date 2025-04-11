package logic

import (
	"github.com/gantoho/osys/internal/middleware"
	"github.com/gantoho/osys/internal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tokenStruct struct {
	Token string `json:"token" form:"token" binding:"required"`
}

func GetUserIDByToken(c *gin.Context) {
	var tokenS tokenStruct
	_ = c.ShouldBind(&tokenS)
	newJwt := middleware.NewJWT()
	ret, err := newJwt.ParseToken(tokenS.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{
			Code:    http.StatusBadRequest,
			Message: "parse token failed",
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code:    http.StatusOK,
		Message: "success",
		Data: gin.H{
			"userId":   ret.UserID,
			"userName": ret.Issuer,
		},
	})
}
