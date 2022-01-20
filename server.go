package main

import (
	"github.com/gin-gonic/gin"
	"yuquey/controller"
)

func main() {
	r := gin.Default()

	//apiV1
	apiV1 := r.Group("/api/V1")
	//ws := r.Group("/ws")

	//test
	test := apiV1.Group("/auth")
	test.POST("/register", controller.Register)

	panic(r.Run("0.0.0.0:8080"))
}
