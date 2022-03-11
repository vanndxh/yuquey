package api

import (
	"github.com/gin-gonic/gin"
	"yuquey/database"
	"yuquey/model"
)

func GetNotice(c *gin.Context) {
	var n model.Notice
	database.DB.Find(&n)
	c.JSON(200, gin.H{"data": n.NoticeContent})
}

func ChangeNotice(c *gin.Context) {
	newNotice := c.PostForm("newNotice")
	var n model.Notice
	database.DB.Model(&n).Update("notice_content", newNotice)
	c.JSON(200, gin.H{"msg": "ok"})
}
