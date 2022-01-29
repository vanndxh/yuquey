package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func SubmitFeedback(c *gin.Context) {
	feedbackInfo := c.PostForm("feedbackInfo")
	if len(feedbackInfo) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "反馈内容不能为空！"})
		return
	}
	newFeedback := model.Feedback{
		FeedbackInfo: feedbackInfo,
	}
	err := database.DB.Create(&newFeedback).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "反馈提交成功！",
	})
}
