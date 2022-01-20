package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/model"
)

func GetUserInfo(c *gin.Context) {
	var user model.User

	// 获取登录信息

	// 查找用户

	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["UserId"] = user.UserId
	returnJSON["Username"] = user.Username
	returnJSON["Password"] = user.Password
	returnJSON["UserInfo"] = user.UserInfo
	returnJSON["ArticleAmount"] = user.ArticleAmount
	returnJSON["LikeTotal"] = user.LikeTotal

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}
