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
		Username: username,
		Password: password,
	}
	err := database.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功！"})
}

func SignIn(c *gin.Context) {
	// 获取参数
	var u model.User
	userId := c.PostForm("userId")
	password := c.PostForm("password")
	// 查找用户是否存在并验证密码
	result := database.DB.Find(&u, "user_id=? AND password=?", userId, password)
	// 返回结果
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "登录成功！"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "账号或密码错误！"})
	}
}

func GetUserInfo(c *gin.Context) {
	// 获取登录信息
	var user model.User
	userId := c.PostForm("userId")
	// 查找用户
	database.DB.Find(&user, "user_id=?", userId)
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["userId"] = user.UserId
	returnJSON["username"] = user.Username
	returnJSON["password"] = user.Password
	returnJSON["userInfo"] = user.UserInfo
	returnJSON["articleAmount"] = user.ArticleAmount
	returnJSON["likeTotal"] = user.LikeTotal
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}

func UpdateUserInfo(c *gin.Context) {
	// 获取数据
	var u model.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	userInfo := c.PostForm("userInfo")
	// 找到对应记录
	result := database.DB.Find(&u, "user_id=?", 1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// update
	database.DB.Model(&u).Update("user_name", username)
	database.DB.Model(&u).Update("password", password)
	database.DB.Model(&u).Update("user_info", userInfo)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "用户个人信息修改成功！"})
}
