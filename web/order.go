package web

import (
	"dormitory_interface/rabbitmq"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderInput struct {
	BuildingNo string `json:"building_no" binding:"required"`
}

func OrderHandler(ctx *gin.Context) {
	var orderInput OrderInput
	value, _ := ctx.Get("JWT_PAYLOAD")
	username := value.(jwt.MapClaims)["username"].(string)
	if err := ctx.Bind(&orderInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rabbitmq.Sendmsg(username + " " + orderInput.BuildingNo)
}
