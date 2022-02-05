package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func DeleteComment(c *gin.Context) {
	// 获取数据
	var cc model.Comment
	articleId, err := strconv.Atoi(c.PostForm("articleId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 删除操作
	result := database.DB.Delete(&cc, "comment_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "评论删除成功！",
	})
}
