package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

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
