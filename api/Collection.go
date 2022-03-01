package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yuquey/database"
	"yuquey/model"
)

func HandleCollection(c *gin.Context) {
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
	handle, err3 := strconv.Atoi(c.PostForm("handle")) // 0-add 1-delete
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	if handle == 0 {
		// 创建新实例
		newStar := model.Collection{
			UserId:    userId,
			ArticleId: articleId,
		}
		err4 := database.DB.Create(&newStar).Error
		if err4 != nil {
			fmt.Println(err4)
			return
		}
	} else {
		var s model.Collection
		// 删除对应实例
		result := database.DB.Delete(&s, "user_id=? AND article_id=?", userId, articleId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
			return
		}
	}
	// 文章收藏数
	var a model.Article
	res := database.DB.Find(&a, "article_id=?", articleId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	starNow := a.CollectionAmount
	if handle == 0 {
		database.DB.Model(&a).Update("collection_amount", starNow+1)
	} else {
		database.DB.Model(&a).Update("collection_amount", starNow-1)
	}
	// 如果是收藏，发消息给作者
	if handle == 0 && a.ArticleAuthor != userId {
		newMessage := model.Message{
			UserId:    a.ArticleAuthor,
			Type:      2,
			Read:      1,
			Op:        userId,
			ArticleId: articleId,
			Time:      time.Now(),
		}
		err5 := database.DB.Create(&newMessage).Error
		if err5 != nil {
			fmt.Println(err5)
			return
		}
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "成功！",
	})
}

func GetFavorite(c *gin.Context) { // 根据用户获取所有收藏的文章
	userId := c.DefaultQuery("userId", "")

	var s []model.Collection
	subQuery := database.DB.Select("article_id").Where("user_id=?", userId).Find(&s)
	if subQuery.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": subQuery.Error.Error()})
		return
	}

	idSlice := make([]int, len(s))
	for i := range s {
		idSlice[i] = s[i].ArticleId
	}
	var as []model.Article
	result := database.DB.Find(&as, idSlice)

	for i := range as {
		var u model.User
		database.DB.Find(&u, "user_id=?", as[i].ArticleAuthor)
		as[i].AuthorName = u.Username
	}
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": as})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "收藏夹是空的！"})
	}
}
func GetIsCollected(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")
	articleId := c.DefaultQuery("articleId", "")

	var s model.Collection
	result := database.DB.Find(&s, "user_id=? AND article_id=?", userId, articleId)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": false})
		return
	}

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": false})
	}
}
