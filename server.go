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
	a.PUT("/transToTrash", api.TransToTrash)
	a.PUT("/updateArticle", api.UpdateArticle)

	// timeline
	tl := apiV1.Group("/timeline")
	tl.GET("/getTimeline", api.GetTimeline)

	// supportCount
	sc := apiV1.Group("/supportCount")
	sc.GET("/getSupportCount", api.GetSupportCount)

	// feedback
	fb := apiV1.Group("/feedback")
	fb.POST("/submitFeedback", api.SubmitFeedback)

	// team
	t := apiV1.Group("/team")
	t.POST("/createTeam", api.CreateTeam)
	t.DELETE("/deleteTeam", api.DeleteTeam)
	t.GET("/getTeam", api.GetTeam)

	// like
	l := apiV1.Group("/like")
	l.POST("/addLike", api.AddLike)
	l.POST("cancelLike", api.CancelLike)

	// star
	s := apiV1.Group("/star")
	s.POST("/addStar", api.AddStar)
	s.POST("cancelStar", api.CancelStar)
	s.GET("/getFavorite", api.GetFavorite)

	panic(r.Run("0.0.0.0:8080"))
}
