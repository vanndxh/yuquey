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
	var s model.Star
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
	// 删除对应实例
	result := database.DB.Delete(&s, "user_id=? AND article_id=?", userId, articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "取消收藏成功！",
	})
}

func GetFavorite(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")

	var s []model.Star
	subQuery := database.DB.Select("article_id").Where("user_id=?", userId).Find(&s)
	if subQuery.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": subQuery.Error.Error()})
		return
	}

	idSlice := make([]int, len(s))
	for i := range s {
		idSlice[i] = s[i].ArticleId
	}
	var a []model.Article
	result := database.DB.Find(&a, idSlice)

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "收藏夹是空的！"})
	}
}

func GetIsStared(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")
	articleId := c.DefaultQuery("articleId", "")

	var s model.Star
	result := database.DB.Find(&s, "user_id=? AND article_id=?", userId, articleId)

	returnJSON1 := make(map[string]interface{})
	returnJSON1["isStared"] = true
	returnJSON2 := make(map[string]interface{})
	returnJSON2["isStared"] = false
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "true"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "false"})
	}
}
