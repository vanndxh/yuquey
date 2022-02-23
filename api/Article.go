package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func GetArticles(c *gin.Context) {
	var a []model.Article
	articleAuthor := c.DefaultQuery("articleAuthor", "")
	isInTrash := c.DefaultQuery("isInTrash", "")
	result := database.DB.Find(&a, "article_author=? AND is_in_trash=?", articleAuthor, isInTrash)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
}

func SearchArticle(c *gin.Context) {
	searchValue := c.DefaultQuery("searchValue", "")

	var a []model.Article
	result := database.DB.Order("hot desc").Find(&a, "article_name like ?", "%"+searchValue+"%")
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 200, "data": "none"})
	}
}

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
	// 给用户文章总数++
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
	var u model.User
	res := database.DB.Find(&u, "user_id=?", a.ArticleAuthor)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	// 返回表单
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a, "authorName": u.Username})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "This article is not exist."})
	}
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
	database.DB.Model(&a).Update("article_name", articleName)
	database.DB.Model(&a).Update("article_content", articleContent)
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

func GetHotArticle(c *gin.Context) {
	// 获取数据
	var a []model.Article
	// 查找对应文章
	result := database.DB.Order("hot desc").Find(&a, "is_in_trash=?", 0)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
}
