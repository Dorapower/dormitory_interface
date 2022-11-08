package main

import (
	"dormitory_interface/auth"
	"dormitory_interface/web"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	r.GET("/hello", web.HelloHandler)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "dormitory middleware",
		Key:         []byte("secret key"),
		IdentityKey: "username",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					"username": v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: auth.Authenticate,
		Authorizator: func(data interface{}, ctx *gin.Context) bool {
			if _, ok := data.(string); ok {
				return true
			} else {
				return false
			}
		},
	})
	if err != nil {
		panic(err)
	}
	r.POST("/login", authMiddleware.LoginHandler)

	authorized := r.Group("/")
	authorized.Use(authMiddleware.MiddlewareFunc())
	{
		authorized.POST("/password", web.PasswordHandler)
		authorized.GET("/student", web.StudentInfoHandler)
		authorized.GET("/building", web.BuildingListHandler)
		authorized.GET("/available", web.AvailableCountHandler)
		authorized.GET("/team/create", web.CreateTeamHandler)
		authorized.POST("/order", web.OrderHandler)
	}
	if err := r.Run(":8090"); err != nil {
		return
	}
}
