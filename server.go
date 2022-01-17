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

	r.POST("/api/auth/register", controller.Register)
	panic(r.Run())
}
