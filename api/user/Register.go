package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func Register(ctx *gin.Context) {
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	rePassword := ctx.PostForm("rePassword")

	// 判断合理性
	if len(password) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能为空！",
		})
		return
	} else if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位！",
		})
		return
	}
	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名不能为空！",
		})
		return
	}
	if rePassword != password {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
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
	ctx.JSON(200, gin.H{
		"msg": "注册成功！",
	})
}
