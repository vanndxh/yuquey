package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

func CreateArticle(c *gin.Context) {
	// 获取数据
	articleName := c.PostForm("articleName")
	articleContent := c.PostForm("articleContent")
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
	}
	err := database.DB.Create(&newArticle).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章创建成功！",
	})
}

func TransToTrash(c *gin.Context) {
	var a model.Article
	articleId := c.PostForm("articleId")
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	database.DB.Model(&a).Update("is_in_trash", 1)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "成功移入垃圾箱！"})
}

func TransOutTrash(c *gin.Context) {
	var a model.Article
	articleId := c.PostForm("articleId")
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	database.DB.Model(&a).Update("is_in_trash", 0)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "成功移出垃圾箱！"})
}

func DeleteArticle(c *gin.Context) {
	// 获取文章id
	var a model.Article
	articleId := c.PostForm("articleId")
	// 删除操作
	result := database.DB.Delete(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章删除成功！",
	})
}

func GetArticleInfo(c *gin.Context) {
	// 获取数据
	var a model.Article
	articleId := c.PostForm("articleId")
	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", articleId)
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount
	returnJSON["Hot"] = a.Hot
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": returnJSON})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "This article is not exist."})
	}
}

func GetHotArticle(c *gin.Context) {
	// 获取数据
	var a model.Article
	articleId := c.PostForm("articleId")
	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", articleId)
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount
	returnJSON["Hot"] = a.Hot
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": returnJSON})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "This article is not exist."})
	}
}

func UpdateArticle(c *gin.Context) {
	// 获取数据
	var a model.Article
	articleName := c.PostForm("articleName")
	articleContent := c.PostForm("articleContent")
	// 判断合理性
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
	// 先找到对应文章数据,即a为要update的记录
	result := database.DB.Find(&a, "article_id=?", 100001)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// update
	database.DB.Model(&a).Update("article_name", articleName)
	database.DB.Model(&a).Update("article_content", articleContent)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "文章修改成功！"})
}
