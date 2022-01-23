package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/model"
)

func GetTimeline(c *gin.Context) {
	var tl model.Timeline
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["Time"] = tl.Time
	returnJSON["Title"] = tl.Title
	returnJSON["Type"] = tl.Type
	returnJSON["Content"] = tl.Content
	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}
