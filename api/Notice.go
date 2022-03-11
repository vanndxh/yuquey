package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"yuquey/database"
	"yuquey/model"
)

func GetNotice(c *gin.Context) {
	var n model.Notice
	database.DB.Find(&n)
	c.JSON(200, gin.H{"data": n})
}

func ChangeNotice(c *gin.Context) {
	newNotice := c.PostForm("newNotice")
	var n model.Notice
	database.DB.Model(&n).Update("notice_content", newNotice)
	database.DB.Model(&n).Update("time", time.Now())
	var u model.User
	database.DB.Model(&u).Update("read_notice", 0)
	c.JSON(200, gin.H{"msg": "ok"})
}
