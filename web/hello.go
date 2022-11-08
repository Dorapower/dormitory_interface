package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world")
}
