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

func HandleLike(c *gin.Context) {
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
	handle, err3 := strconv.Atoi(c.PostForm("handle"))
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	if handle == 0 {
		newLike := model.Like{
			UserId:    userId,
			ArticleId: articleId,
		}
		err4 := database.DB.Create(&newLike).Error
		if err4 != nil {
			fmt.Println(err4)
			return
		}
	} else {
		// 寻找对应实例
		var l model.Like
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
	}
	// 找到对应记录
	var a model.Article
	res := database.DB.Find(&a, "article_id=?", articleId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	var u model.User
	result := database.DB.Find(&u, "user_id=?", a.ArticleAuthor)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	// 给用户点赞总数
	likeNow := u.LikeAmount
	if handle == 0 {
		database.DB.Model(&u).Update("like_amount", likeNow+1)
	} else {
		database.DB.Model(&u).Update("like_amount", likeNow-1)
	}

	// 给文章点赞数
	likeNow2 := a.LikeAmount
	if handle == 0 {
		database.DB.Model(&a).Update("like_amount", likeNow2+1)
	} else {
		database.DB.Model(&a).Update("like_amount", likeNow2-1)
	}

	// 给文章热度
	hotNow := a.Hot
	if handle == 0 {
		database.DB.Model(&a).Update("hot", hotNow+1)
	} else {
		database.DB.Model(&a).Update("hot", hotNow-1)
	}

	// 如果是点赞，发消息给作者
	if handle == 0 && a.ArticleAuthor != userId {
		newMessage := model.Message{
			UserId:    a.ArticleAuthor,
			Type:      1,
			Op:        userId,
			ArticleId: articleId,
			Read:      1,
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

func GetIsLiked(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")
	articleId := c.DefaultQuery("articleId", "")

	var l model.Like
	result := database.DB.Find(&l, "user_id=? AND article_id=?", userId, articleId)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": false, "msg": "no like"})
		return
	}

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": false})
	}
}
