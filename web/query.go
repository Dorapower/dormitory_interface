package web

import (
	"dormitory_interface/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentInfoHandler(ctx *gin.Context) {
	value, _ := ctx.Get("JWT_PAYLOAD")
	username := value.(jwt.MapClaims)["username"].(string)
	data := sql.QueryStudent(username)
	ctx.JSON(http.StatusOK, data)
}

func BuildingListHandler(ctx *gin.Context) {
	value, _ := ctx.Get("JWT_PAYLOAD")
	username := value.(jwt.MapClaims)["username"].(string)
	data := sql.QueryBuildingList(username)
	ctx.JSON(http.StatusOK, data)
}

func AvailableCountHandler(ctx *gin.Context) {
	buildingNo := ctx.Query("building")
	data := sql.QueryAvailableCount(buildingNo)
	ctx.JSON(http.StatusOK, data)
}
