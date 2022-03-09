package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
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
		UserInfo: "暂无~",
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

func DeleteUser(c *gin.Context) {
	userId := c.Query("userId")

	// 删除用户
	var u model.User
	database.DB.Delete(&u, "user_id=?", userId)

	// 清除他的所有文章
	var a model.Article
	database.DB.Delete(&a, "article_author=?", userId)

	// 退出他在的所有小组
	// here?

	// 清楚所有评论
	var cc model.Comment
	database.DB.Delete(&cc, "user_id=?", userId)

	// 清除所有通知
	var m model.Message
	database.DB.Delete(&m, "user_id=?", userId)

	// 清除关注信息
	var f1 model.Follow
	database.DB.Delete(&f1, "up_id=?", userId)
	var f2 model.Follow
	database.DB.Delete(&f2, "follower_id=?", userId)

	// 清除点赞收藏信息
	var l model.Like
	database.DB.Delete(&l, "user_id=?", userId)
	var co model.Collection
	database.DB.Delete(&co, "user_id=?", userId)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "ok"})
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
	database.DB.Model(&u).Update("username", username)
	database.DB.Model(&u).Update("password", password)
	database.DB.Model(&u).Update("user_info", userInfo)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "用户个人信息修改成功！"})
}
func RenewVip(c *gin.Context) {
	// 获取参数
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	code := c.PostForm("code")
	// 找到user
	var u model.User
	database.DB.Find(&u, "user_id=?", userId)
	// 准备各种时间段的duration
	tp, _ := time.ParseDuration("720h")
	// 判定code进行操作
	if code == "xiaoheihaoshuai" {
		if u.Vip.Before(time.Now()) {
			newVip := time.Now().Add(tp)
			res := database.DB.Model(&u).Where("user_id=?", userId).Update("vip", newVip)
			if res.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
				return
			}
		} else {
			newVip := u.Vip.Add(tp)
			res := database.DB.Model(&u).Where("user_id=?", userId).Update("vip", newVip)
			if res.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
				return
			}
		}
		c.JSON(200, gin.H{"msg": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
	}
}

func GetUserInfo(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")

	var user model.User
	database.DB.Find(&user, "user_id=?", userId)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": user})
}
func GetAllUsers(c *gin.Context) {
	var us []model.User
	database.DB.Order("user_id").Find(&us)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": us})
}
