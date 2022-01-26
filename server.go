package main

import (
	"github.com/gin-gonic/gin"
	"yuquey/api"
)

func main() {
	r := gin.Default()

	//apiV1
	apiV1 := r.Group("/api/V1")
	//ws := r.Group("/ws")

	// user
	us := apiV1.Group("/user")
	us.POST("/register", api.Register)
	us.POST("/signIn", api.SignIn)
	us.GET("/getUserInfo", api.GetUserInfo)

	// article
	a := apiV1.Group("/article")
	a.GET("/getArticleInfo", api.GetArticleInfo)
	a.POST("/createArticle", api.CreateArticle)
	a.GET("/getHotArticle", api.GetHotArticle)
	a.DELETE("/deleteArticle", api.DeleteArticle)
	a.PUT("/updateArticle", api.UpdateArticle)

	// timeline
	tl := apiV1.Group("/timeline")
	tl.GET("/getTimeline", api.GetTimeline)

	// supportCount
	sc := apiV1.Group("/supportCount")
	sc.GET("/getSupportCount", api.GetSupportCount)

	panic(r.Run("0.0.0.0:8080"))
}
