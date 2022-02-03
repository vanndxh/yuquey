package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func CreateComment(c *gin.Context) {
	// 获取数据
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	articleId, err2 := strconv.Atoi(c.PostForm("articleId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	commentContent := c.PostForm("commentContent")
	// 创建新实例
	newComment := model.Comment{
		UserId:         userId,
		ArticleId:      articleId,
		CommentContent: commentContent,
	}
	err3 := database.DB.Create(&newComment).Error
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "评论成功！",
	})
}
