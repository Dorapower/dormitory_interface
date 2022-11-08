package web

import (
	"dormitory_interface/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTeamHandler(ctx *gin.Context) {
	value, _ := ctx.Get("JWT_PAYLOAD")
	username := value.(jwt.MapClaims)["username"].(string)
	if sql.QueryStudentIncluded(username) != -1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "student already in a team"})
		return
	}
	sql.CreateNewTeam(username)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
