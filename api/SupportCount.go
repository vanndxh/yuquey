package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func GetSupportCount(c *gin.Context) {
	// 获取数据
	var sc model.SupportCount
	database.DB.Find(&sc)
	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": sc.Count})
	// 访问一次数字加一
	var newCount = sc.Count + 1
	database.DB.Model(&sc).Update("count", newCount)
}
