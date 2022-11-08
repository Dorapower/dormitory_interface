package web

import (
	"dormitory_interface/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordInput struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}

func testValid(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func PasswordHandler(ctx *gin.Context) {
	value, _ := ctx.Get("JWT_PAYLOAD")
	username := value.(jwt.MapClaims)["username"].(string)

	var passwordInput PasswordInput
	if err := ctx.ShouldBind(&passwordInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !sql.QueryLogin(username, passwordInput.OldPassword) {
		ctx.JSON(400, gin.H{"error": "Wrong password"})
		return
	}
	if !testValid(passwordInput.NewPassword) {
		ctx.JSON(400, gin.H{"error": "Invalid password"})
		return
	}
	sql.UpdatePassword(username, passwordInput.NewPassword)
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed"})
}
