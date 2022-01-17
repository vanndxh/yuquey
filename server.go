package main

import (
	"github.com/gin-gonic/gin"
	"yuquey/common"
	"yuquey/controller"
)

func main() {
	r := gin.Default()
	db := common.InitDB()
	defer db.Close()

	//apiV1
	apiV1 := r.Group("/api/v1")
	//ws := r.Group("/ws")

	//test
	test := apiV1.Group("/auth")
	test.POST("/register", controller.Register)

	panic(r.Run())
}
