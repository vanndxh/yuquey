package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func SubmitFeedback(c *gin.Context) {
	feedbackInfo := c.PostForm("feedbackInfo")
	if len(feedbackInfo) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "反馈内容不能为空！"})
		return
	}
	userId, _ := strconv.Atoi(c.PostForm("userId"))

	newFeedback := model.Feedback{
		FeedbackInfo: feedbackInfo,
		UserId:       userId,
	}
	err := database.DB.Create(&newFeedback).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"msg": "反馈提交成功！"})
}

func DeleteFeedback(c *gin.Context) {
	id := c.Query("feedbackId")
	var f model.Feedback
	res := database.DB.Delete(&f, "feedback_id=?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}

func GetFeedbacks(c *gin.Context) {
	var fs []model.Feedback
	database.DB.Find(&fs)
	c.JSON(200, gin.H{"data": fs})
}
