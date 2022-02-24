package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func GetMessages(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")
	var ms []model.Message
	res := database.DB.Order("time desc").Find(&ms, "user_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range ms {
		var u model.User
		database.DB.Find(&u, "user_id=?", ms[i].UserId)
		ms[i].Username = u.Username
		if ms[i].Type == 0 {
			ms[i].TypeName = "点赞"
		} else if ms[i].Type == 1 {
			ms[i].TypeName = "收藏"
		} else {
			ms[i].TypeName = "评论"
		}
		var op model.User
		database.DB.Find(&op, "user_id=?", ms[i].Op)
		ms[i].OpName = op.Username
		var a model.Article
		database.DB.Find(&a, "article_id=?", ms[i].ArticleId)
		ms[i].ArticleName = a.ArticleName
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ms})
}

func ReadMessage(c *gin.Context) {
	messageId, err := strconv.Atoi(c.PostForm("messageId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var m model.Message
	database.DB.Model(&m).Where("message_id=?", messageId).Update("read", 1)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "ok"})
}

func ReadAllMessages(c *gin.Context) {
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var ms []model.Message
	res := database.DB.Find(&ms, "user_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range ms {
		if ms[i].Read == 0 {
			database.DB.Model(&ms[i]).Update("read", 1)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "ok"})
}

func GetRead(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")
	var ms []model.Message
	res := database.DB.Find(&ms, "user_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range ms {
		if ms[i].Read == 0 {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": false})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": true})
}
