package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/model"
)

func GetTimeline(c *gin.Context) {
	var tl model.Timeline
	fmt.Println(tl)

	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "test"})

}
