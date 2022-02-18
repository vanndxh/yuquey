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
	V1user.POST("/updateUserInfo", api.UpdateUserInfo)
	V1user.GET("/getUserInfo", api.GetUserInfo)
	V1user.POST("/login", api.Login)

	// article
	V1article := apiV1.Group("/article")
	V1article.GET("/getArticles", api.GetArticles)
	V1article.GET("/getArticleInfo", api.GetArticleInfo)
	V1article.POST("/createArticle", api.CreateArticle)
	V1article.DELETE("/deleteArticle", api.DeleteArticle)
	V1article.POST("/updateArticle", api.UpdateArticle)
	V1article.POST("/transTrash", api.TransTrash)
	V1article.GET("/searchArticle", api.SearchArticle)
	V1article.GET("/getHotArticle", api.GetHotArticle)

	// team
	V1team := apiV1.Group("/team")
	V1team.GET("/getTeams", api.GetTeams)
	V1team.POST("/createTeam", api.CreateTeam)
	V1team.DELETE("/deleteTeam", api.DeleteTeam)
	V1team.POST("/updateTeamInfo", api.UpdateTeamInfo)
	V1team.GET("/getTeamInfo", api.GetTeamInfo)
	V1team.POST("/addTeamUser", api.AddTeamUser)
	V1team.POST("/punch", api.Punch)
	V1team.POST("/quitTeam", api.QuitTeam)
	V1team.GET("/getTeamArticles", api.GetTeamArticles)
	V1team.GET("/getTeamMembers", api.GetTeamMembers)

	// comment
	V1comment := apiV1.Group("/comment")
	V1comment.POST("/createComment", api.CreateComment)
	V1comment.DELETE("/deleteComment", api.DeleteComment)

	// like
	V1like := apiV1.Group("/like")
	V1like.POST("/addLike", api.AddLike)
	V1like.DELETE("cancelLike", api.CancelLike)
	V1like.GET("/getIsLiked", api.GetIsLiked)

	// star
	V1star := apiV1.Group("/star")
	V1star.POST("/addStar", api.AddStar)
	V1star.POST("cancelStar", api.CancelStar)
	V1star.GET("/getFavorite", api.GetFavorite)
	V1star.GET("/getIsStared", api.GetIsStared)

	// timeline
	V1timeline := apiV1.Group("/timeline")
	V1timeline.GET("/getTimeline", api.GetTimeline)

	// supportCount
	V1supportCount := apiV1.Group("/supportCount")
	V1supportCount.GET("/getSupportCount", api.GetSupportCount)

	// feedback
	V1feedback := apiV1.Group("/feedback")
	V1feedback.POST("/submitFeedback", api.SubmitFeedback)

	panic(r.Run("0.0.0.0:8088"))
}
