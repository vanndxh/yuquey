package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func AddStar(c *gin.Context) {
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
	newStar := model.Star{
		UserId:    userId,
		ArticleId: articleId,
	}
	err3 := database.DB.Create(&newStar).Error
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "收藏成功！",
	})
}

func CancelStar(c *gin.Context) {
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
	result2 := database.DB.Delete(&l, "user_id=? AND article_id=?", userId, articleId)
	if result2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "取消收藏成功！",
	})
}

func GetFavorite(c *gin.Context) {
	// 获取数据
	var s []model.Star
	var a []model.Article
	userId := c.PostForm("userId")
	// 先取articleId
	subQuery := database.DB.Select("article_id").Where("user_id=?", userId).Find(&s)
	if subQuery.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": subQuery.Error.Error()})
		return
	}
	// id切片存入[]int
	idSlice := make([]int, len(s))
	for i := range s {
		idSlice[i] = s[i].ArticleId
	}
	result := database.DB.Find(&a, idSlice)
	// 返回结果
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "收藏夹是空的！"})
	}
}
