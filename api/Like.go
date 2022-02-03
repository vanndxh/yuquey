package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func AddLike(c *gin.Context) {
	// 获取数据
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	articleId, err2 := strconv.Atoi(c.PostForm("articleId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 创建新实例
	newLike := model.Like{
		UserId:    userId,
		ArticleId: articleId,
	}
	err3 := database.DB.Create(&newLike).Error
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "点赞成功！",
	})
}

func CancelLike(c *gin.Context) {
	var l model.Like
	// 获取数据
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	articleId, err2 := strconv.Atoi(c.PostForm("articleId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 寻找对应实例
	result := database.DB.Find(&l, "user_id=? AND article_id=?", userId, articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 数据库中删除这条记录
	result2 := database.DB.Delete(&l, "user_id=? AND article_id=?", userId, articleId)
	if result2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "取消点赞成功！",
	})
}

func JudgeIsLiked(c *gin.Context) {
	var l model.Like
	// 获取数据
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	articleId, err2 := strconv.Atoi(c.PostForm("articleId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 寻找对应实例
	result := database.DB.Find(&l, "user_id=? AND article_id=?", userId, articleId)
	// 返回表单
	returnJSON1 := make(map[string]interface{})
	returnJSON1["isLiked"] = true
	returnJSON2 := make(map[string]interface{})
	returnJSON2["isLiked"] = false
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": returnJSON1})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": returnJSON2})
	}
}
