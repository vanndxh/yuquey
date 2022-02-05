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
	V1user := apiV1.Group("/user")
	V1user.POST("/register", api.Register)
	V1user.POST("/signIn", api.SignIn)
	V1user.GET("/getUserInfo", api.GetUserInfo)
	V1user.PUT("/updateUserInfo", api.UpdateUserInfo)

	// article
	V1article := apiV1.Group("/article")
	V1article.GET("/getArticleInfo", api.GetArticleInfo)
	V1article.POST("/createArticle", api.CreateArticle)
	V1article.GET("/getHotArticle", api.GetHotArticle)
	V1article.DELETE("/deleteArticle", api.DeleteArticle)
	V1article.PUT("/transToTrash", api.TransToTrash)
	V1article.PUT("/transOutTrash", api.TransOutTrash)
	V1article.PUT("/updateArticle", api.UpdateArticle)
	V1article.GET("/searchArticle", api.SearchArticle)

	// team
	V1team := apiV1.Group("/team")
	V1team.POST("/createTeam", api.CreateTeam)
	V1team.DELETE("/deleteTeam", api.DeleteTeam)
	V1team.GET("/getTeams", api.GetTeams)
	V1team.GET("/getTeamInfo", api.GetTeamInfo)
	V1team.PUT("/updateTeamInfo", api.UpdateTeamInfo)
	V1team.PUT("/addTeamUser", api.AddTeamUser)

	// comment
	V1comment := apiV1.Group("/comment")
	V1comment.POST("/createComment", api.CreateComment)
	V1comment.DELETE("/deleteComment", api.DeleteComment)

	// like
	V1like := apiV1.Group("/like")
	V1like.POST("/addLike", api.AddLike)
	V1like.POST("cancelLike", api.CancelLike)
	V1like.GET("/judgeIsLiked", api.JudgeIsLiked)

	// star
	V1star := apiV1.Group("/star")
	V1star.POST("/addStar", api.AddStar)
	V1star.POST("cancelStar", api.CancelStar)
	V1star.GET("/getFavorite", api.GetFavorite)
	V1like.GET("/judgeIsStared", api.JudgeIsStared)

	// timeline
	V1timeline := apiV1.Group("/timeline")
	V1timeline.GET("/getTimeline", api.GetTimeline)

	// supportCount
	V1supportCount := apiV1.Group("/supportCount")
	V1supportCount.GET("/getSupportCount", api.GetSupportCount)

	// feedback
	V1feedback := apiV1.Group("/feedback")
	V1feedback.POST("/submitFeedback", api.SubmitFeedback)

	panic(r.Run("0.0.0.0:8080"))
}
