package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/model"
)

func GetSupportCount(c *gin.Context) {
	var sc model.SupportCount
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["Count"] = sc.Count

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}
