package main

import (
	"github.com/gin-gonic/gin"
	"yuquey/api"
	"yuquey/database"
	"yuquey/model"
)

func main() {
	database.DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Timeline{}, &model.SupportCount{}, &model.Like{},
		&model.Collection{}, &model.Feedback{}, &model.Team{}, &model.Comment{}, &model.Follow{}, &model.Message{},
		&model.TeamUser{}, &model.Notice{})

	r := gin.Default()

	//apiV1
	apiV1 := r.Group("/api/V1")
	apiV1.GET("/search", api.Search)
	//ws := r.Group("/ws")

	// user
	V1user := apiV1.Group("/user")
	V1user.POST("/register", api.Register)
	V1user.POST("/updateUserInfo", api.UpdateUserInfo)
	V1user.GET("/getUserInfo", api.GetUserInfo)
	V1user.POST("/login", api.Login)
	V1user.GET("/getAllUsers", api.GetAllUsers)
	V1user.DELETE("/deleteUser", api.DeleteUser)
	V1user.POST("/renewVip", api.RenewVip)
	V1user.POST("/addAuth", api.AddAuth)
	V1user.POST("/readNotice", api.ReadNotice)

	// article
	V1article := apiV1.Group("/article")
	V1article.GET("/getArticles", api.GetArticles)
	V1article.GET("/getArticleInfo", api.GetArticleInfo)
	V1article.POST("/createArticle", api.CreateArticle)
	V1article.DELETE("/deleteArticle", api.DeleteArticle)
	V1article.POST("/updateArticle", api.UpdateArticle)
	V1article.POST("/transTrash", api.TransTrash)
	V1article.GET("/getHotArticles", api.GetHotArticles)
	V1article.GET("/getAllArticles", api.GetAllArticles)
	V1article.GET("/getFollowArticles", api.GetFollowArticles)
	V1article.DELETE("/deleteAllArticle", api.DeleteAllArticle)

	// team
	V1team := apiV1.Group("/team")
	V1team.GET("/getTeams", api.GetTeams)
	V1team.POST("/createTeam", api.CreateTeam)
	V1team.DELETE("/deleteTeam", api.DeleteTeam)
	V1team.POST("/updateTeamInfo", api.UpdateTeamInfo)
	V1team.GET("/getTeamInfo", api.GetTeamInfo)
	V1team.POST("/handleTeamUser", api.HandleTeamUser)
	V1team.POST("/punch", api.Punch)
	V1team.POST("/quitTeam", api.QuitTeam)
	V1team.GET("/getTeamArticles", api.GetTeamArticles)
	V1team.GET("/getTeamMembers", api.GetTeamMembers)
	V1team.GET("/getAllTeams", api.GetAllTeams)
	V1team.POST("/handleManager", api.HandleManager)

	// comment
	V1comment := apiV1.Group("/comment")
	V1comment.POST("/createComment", api.CreateComment)
	V1comment.DELETE("/deleteComment", api.DeleteComment)
	V1comment.GET("/getArticleComment", api.GetArticleComment)
	V1comment.GET("/getUserComment", api.GetUserComment)
	V1comment.GET("/getAllComments", api.GetAllComments)

	// like
	V1like := apiV1.Group("/like")
	V1like.POST("/handleLike", api.HandleLike)
	V1like.GET("/getIsLiked", api.GetIsLiked)

	// collection
	V1star := apiV1.Group("/collection")
	V1star.POST("/handleCollection", api.HandleCollection)
	V1star.GET("/getFavorite", api.GetFavorite)
	V1star.GET("/getIsCollected", api.GetIsCollected)

	// message
	V1message := apiV1.Group("/message")
	V1message.GET("/getMessages", api.GetMessages)
	V1message.POST("/readMessage", api.ReadMessage)
	V1message.POST("/readAllMessages", api.ReadAllMessages)
	V1message.GET("/getRead", api.GetRead)
	V1message.POST("/sendMessage", api.SendMessage)

	// timeline
	V1timeline := apiV1.Group("/timeline")
	V1timeline.GET("/getTimeline", api.GetTimeline)
	V1timeline.POST("/addTimeline", api.AddTimeline)
	V1timeline.DELETE("/deleteTimeline", api.DeleteTimeline)

	// follow
	V1follow := apiV1.Group("/follow")
	V1follow.POST("/handleFollow", api.HandleFollow)
	V1follow.GET("/getFollows", api.GetFollows)
	V1follow.GET("/getIsFollowed", api.GetIsFollowed)

	// supportCount
	V1supportCount := apiV1.Group("/supportCount")
	V1supportCount.GET("/getSupportCount", api.GetSupportCount)

	// notice
	V1notice := apiV1.Group("/notice")
	V1notice.GET("/getNotice", api.GetNotice)
	V1notice.POST("/changeNotice", api.ChangeNotice)

	// feedback
	V1feedback := apiV1.Group("/feedback")
	V1feedback.POST("/submitFeedback", api.SubmitFeedback)
	V1feedback.GET("/getFeedbacks", api.GetFeedbacks)
	V1feedback.DELETE("/deleteFeedback", api.DeleteFeedback)

	panic(r.Run("0.0.0.0:8088"))
}
