package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yuquey/database"
	"yuquey/model"
)

func SendMessage(c *gin.Context) {
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	messageContent := c.PostForm("messageContent")

	newMessage := model.Message{
		Content: messageContent,
		Read:    1,
		Time:    time.Now(),
		UserId:  userId,
		Type:    5,
	}

	res := database.DB.Create(&newMessage).Error
	if res != nil {
		fmt.Println(res)
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}

func ReadMessage(c *gin.Context) {
	messageId, err := strconv.Atoi(c.PostForm("messageId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var m model.Message
	database.DB.Model(&m).Where("message_id=?", messageId).Update("read", 2)
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
		if ms[i].Read == 1 {
			database.DB.Model(&ms[i]).Update("read", 2)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "ok"})
}

func GetRead(c *gin.Context) { // 判断是否有未读消息
	userId := c.DefaultQuery("userId", "")

	var ms []model.Message
	res := database.DB.Find(&ms, "user_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range ms {
		if ms[i].Read == 1 {
			c.JSON(http.StatusOK, gin.H{"data": false})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
func GetMessages(c *gin.Context) { // 根据用户获取消息
	userId := c.DefaultQuery("userId", "")
	handle1 := c.DefaultQuery("handle1", "0")
	handle2 := c.DefaultQuery("handle2", "0")

	var ms []model.Message
	// 选择筛选模式
	// handle1: 0-all 1-no 2-yes
	// handle2: 0-all 1-like 2-collection 3-comment 4-follow 5-other
	if handle1 == "0" && handle2 == "0" {
		res := database.DB.Order("time desc").Find(&ms, "user_id=?", userId)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	} else if handle1 == "0" && handle2 != "0" {
		res := database.DB.Order("time desc").Find(&ms, "user_id=? AND type=?", userId, handle2)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	} else if handle1 != "0" && handle2 == "0" {
		res := database.DB.Order("time desc").Find(&ms, "user_id=? AND read=?", userId, handle1)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	} else {
		res := database.DB.Order("time desc").Find(&ms, "user_id=? AND read=? AND type=?", userId, handle1, handle2)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	}
	// 生成真正的Content
	for i := range ms {
		var typeName string
		var op model.User
		database.DB.Find(&op, "user_id=?", ms[i].Op)
		if ms[i].Type == 1 {
			typeName = "点赞"
		} else if ms[i].Type == 2 {
			typeName = "收藏"
		} else if ms[i].Type == 3 {
			typeName = "评论"
		} else if ms[i].Type == 4 {
			typeName = "关注"
		}

		if ms[i].Type == 4 {
			ms[i].Content = "用户" + op.Username + typeName + "了您"
		} else if ms[i].Type == 1 || ms[i].Type == 2 || ms[i].Type == 3 {
			var a model.Article
			database.DB.Find(&a, "article_id=?", ms[i].ArticleId)
			ms[i].Content = "用户" + op.Username + typeName + "了您的文章《" + a.ArticleName + "》"
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ms})
}
