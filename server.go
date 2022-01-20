package main

import (
	"github.com/gin-gonic/gin"
	"yuquey/api/user"
)

func main() {
	r := gin.Default()

	//apiV1
	apiV1 := r.Group("/api/V1")
	//ws := r.Group("/ws")

	// user
	us := apiV1.Group("/auth")
	us.POST("/register", user.Register)
	us.GET("getUserInfo", user.GetUserInfo)

	panic(r.Run("0.0.0.0:8080"))
}
