package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yuquey/database"
	"yuquey/model"
	"yuquey/util"
)

func CreateArticle(c *gin.Context) {
	// 获取数据
	articleContent := c.PostForm("articleContent")
	articleName := c.PostForm("articleName")
	articleAuthor, err := strconv.Atoi(c.PostForm("articleAuthor"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断数据合理性
	if len(articleName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "文章名称不能为空！",
		})
		return
	}
	if len(articleContent) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "文章内容不能为空！",
		})
		return
	}
	// 创建新文章
	newArticle := model.Article{
		ArticleName:    articleName,
		ArticleContent: articleContent,
		ArticleAuthor:  articleAuthor,
		Time:           time.Now(),
	}
	err2 := database.DB.Create(&newArticle).Error
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 给用户文章总数++
	var u model.User
	result := database.DB.Find(&u, "user_id=?", 1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	articleNow := u.ArticleAmount
	database.DB.Model(&u).Update("article_amount", articleNow+1)
	// 返回结果
	c.JSON(200, gin.H{
		"msg":       "文章创建成功！",
		"articleId": newArticle.ArticleId,
	})
}

func DeleteArticle(c *gin.Context) {
	// 获取文章id
	var a model.Article
	articleId := c.DefaultQuery("articleId", "")
	userId := c.DefaultQuery("userId", "")
	// 删除操作
	result := database.DB.Delete(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 给用户文章总数--
	var u model.User
	result2 := database.DB.Find(&u, "user_id=?", userId)
	if result2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	articleNow := u.ArticleAmount
	database.DB.Model(&u).Update("article_amount", articleNow-1)
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章删除成功！",
	})
}

func UpdateArticle(c *gin.Context) {
	// 获取数据
	var a model.Article
	articleId, err := strconv.Atoi(c.PostForm("articleId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	articleName := c.PostForm("newArticleName")
	articleContent := c.PostForm("newArticleContent")
	secret, _ := strconv.Atoi(c.PostForm("newSecret"))

	// 判断合理性
	if len(articleName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "文章名称不能为空！",
		})
		return
	}
	// 先找到对应文章数据
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// update
	database.DB.Model(&a).Where("article_id=?", articleId).Update("article_name", articleName)
	database.DB.Model(&a).Where("article_id=?", articleId).Update("article_content", articleContent)
	database.DB.Model(&a).Where("article_id=?", articleId).Update("secret", secret)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "文章修改成功！"})
}
func TransTrash(c *gin.Context) {
	var a model.Article
	articleId := c.PostForm("articleId")
	handle, err := strconv.Atoi(c.PostForm("handle"))
	if err != nil {
		fmt.Println(err)
		return
	}
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	database.DB.Model(&a).Update("is_in_trash", handle)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "成功移入垃圾箱！"})
}

func GetArticles(c *gin.Context) { // 根据用户获取
	articleAuthor := c.DefaultQuery("articleAuthor", "")
	isInTrash := c.DefaultQuery("isInTrash", "")

	var as []model.Article
	result := database.DB.Find(&as, "article_author=? AND is_in_trash=?", articleAuthor, isInTrash)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	for i := range as {
		var u model.User
		database.DB.Find(&u, "user_id=?", as[i].ArticleAuthor)
		as[i].AuthorName = u.Username
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": as})
}
func GetArticleInfo(c *gin.Context) {
	// 获取数据
	var a model.Article
	articleId := c.DefaultQuery("articleId", "")
	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 获取作者name
	var u model.User
	res := database.DB.Find(&u, "user_id=?", a.ArticleAuthor)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	// 浏览量
	viewNow := a.ViewAmount
	database.DB.Model(&a).Where("article_id=?", articleId).Update("view_amount", viewNow+1)
	// 返回表单
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a, "authorName": u.Username})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "This article is not exist."})
	}
}
func GetHotArticle(c *gin.Context) {
	// 计算最新hot
	util.CalculateHot()
	// 查找对应文章
	var ass []model.Article
	result := database.DB.Where("secret=?", 0).Order("hot desc").Find(&ass, "is_in_trash=?", 0)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ass})
}
func GetAllArticles(c *gin.Context) {
	var as []model.Article
	database.DB.Order("article_id").Find(&as)
	for i := range as {
		var u model.User
		database.DB.Find(&u, "user_id=?", as[i].ArticleAuthor)
		as[i].AuthorName = u.Username
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": as})
}
