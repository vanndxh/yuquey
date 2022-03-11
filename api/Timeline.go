package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func AddTimeline(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	time := c.PostForm("time")
	t := c.PostForm("type")

	newTimeline := model.Timeline{
		Time:    time,
		Title:   title,
		Content: content,
		Type:    t,
	}
	err := database.DB.Create(&newTimeline).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"msg": "ok！"})
}

func DeleteTimeline(c *gin.Context) {
	title := c.Query("title")
	var t model.Timeline
	res := database.DB.Delete(&t, "title=?", title)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}

func GetTimeline(c *gin.Context) {
	// 获取数据
	var tls []model.Timeline
	// 查找全部
	res := database.DB.Order("time desc").Find(&tls)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	// 返回全部
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": tls})
}
