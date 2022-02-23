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

func CreateComment(c *gin.Context) {
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
	commentContent := c.PostForm("commentContent")

	var u model.User
	database.DB.Find(&u, "user_id=?", userId)

	newComment := model.Comment{
		UserId:         userId,
		ArticleId:      articleId,
		CommentContent: commentContent,
	}
	err3 := database.DB.Create(&newComment).Error
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	var a model.Article
	res := database.DB.Find(&a, "article_id=?", articleId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	hotNow := a.Hot
	database.DB.Model(&a).Update("hot", hotNow+2)

	if a.ArticleAuthor != userId {
		newMessage := model.Message{
			UserId:    a.ArticleAuthor,
			Type:      2,
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

	c.JSON(200, gin.H{
		"msg": "评论成功！",
	})
}

func DeleteComment(c *gin.Context) {
	// 获取数据
	var cc model.Comment
	commentId, err := strconv.Atoi(c.PostForm("commentId"))
	if err != nil {
		fmt.Println(err)
		return
	}

	res := database.DB.Find(&cc, "comment_id=?", commentId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	var a model.Article
	res2 := database.DB.Find(&a, "article_id=?", cc.ArticleId)
	if res2.Error != nil {
		fmt.Println(res2.Error)
		return
	}
	hotNow := a.Hot
	database.DB.Model(&a).Update("hot", hotNow-2)

	// 删除操作
	result := database.DB.Delete(&cc, "comment_id=?", commentId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	// 返回结果
	c.JSON(200, gin.H{
		"msg": "评论删除成功！",
	})
}

func GetArticleComment(c *gin.Context) {
	articleId := c.DefaultQuery("articleId", "")

	var cc []model.Comment
	res := database.DB.Find(&cc, "article_id=?", articleId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	//var i model.Comment
	//for i = range cc {
	//	uid := i.UserId
	//	var u model.User
	//	database.DB.Find(&u, "user_id=?", uid)
	//	// 给每个i加上找出来的u.Username
	//	fmt.Println(i)
	//}

	c.JSON(200, gin.H{"status": 200, "data": cc})
}

func GetUserComment(c *gin.Context) {
	userId := c.DefaultQuery("userId", "")

	var cc []model.Comment
	res := database.DB.Find(&cc, "user_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	c.JSON(200, gin.H{"status": 200, "data": cc})
}
