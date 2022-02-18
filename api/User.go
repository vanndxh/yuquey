package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func Register(c *gin.Context) {
	// 获取参数
	username, ok := c.GetPostForm("username")
	if !ok {
		fmt.Println(ok, username, "1")
		return
	}
	password, ok2 := c.GetPostForm("password")
	if !ok2 {
		fmt.Println(ok2, password, "2")
		return
	}
	rePassword, ok3 := c.GetPostForm("rePassword")
	if !ok3 {
		fmt.Println(ok3, rePassword, "3")
		return
	}
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
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "注册成功！", "userId": newUser.UserId, "password": newUser.Password})
}

func Login(c *gin.Context) {
	// 获取参数
	var u model.User
	userId, ok := c.GetPostForm("userId")
	if !ok {
		fmt.Println(ok)
		return
	}
	password, ok2 := c.GetPostForm("password")
	if !ok2 {
		fmt.Println(ok2)
		return
	}
	// 查找用户是否存在并验证密码
	result := database.DB.Find(&u, "user_id=? AND password=?", userId, password)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "账号或密码错误！"})
		return
	}
	// 返回结果
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "登录成功！"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "账号或密码错误！"})
	}
}

func GetUserInfo(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")

	var user model.User
	database.DB.Find(&user, "user_id=?", userId)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": user})
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
