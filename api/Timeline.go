package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func GetTimeline(c *gin.Context) {
	// 获取数据
	var tls []model.Timeline
	// 查找全部
	database.DB.Find(&tls)
	// 返回全部
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": tls})
}
