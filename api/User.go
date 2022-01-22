package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func Register(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")

	// 判断合理性
	if len(password) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能为空！",
		})
		return
	} else if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位！",
		})
		return
	}
	if len(username) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名不能为空！",
		})
		return
	}
	if rePassword != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "两次输入密码需一致！",
		})
		return
	}

	// 创建用户
	newUser := model.User{
		UserId:   "100001",
		Username: username,
		Password: password,
		UserInfo: "",
	}
	err := database.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
	}

	//返回结果
	c.JSON(200, gin.H{
		"msg": "注册成功！",
	})
}

func SignIn(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 判断合理性
	if len(password) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能为空！",
		})
		return
	} else if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位！",
		})
		return
	}
	if len(username) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名不能为空！",
		})
		return
	}

	// 查找用户是否存在并验证密码

	// 返回结果
	c.JSON(200, gin.H{
		"msg": "登录成功！",
	})
}

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
