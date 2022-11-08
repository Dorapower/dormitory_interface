package auth

import (
	"dormitory_interface/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) (user interface{}, err error) {
	var loginInput LoginInput
	if err := ctx.ShouldBind(&loginInput); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	username := loginInput.Username
	password := loginInput.Password
	if sql.QueryLogin(username, password) {
		return username, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
