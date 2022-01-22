package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yuquey/database"
	"yuquey/model"
)

// CreateArticle 新建文章
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

	//创建文章
	newArticle := model.Article{
		ArticleName:    articleName,
		ArticleContent: articleContent,
	}
	err := database.DB.Create(&newArticle).Error
	if err != nil {
		log.Println(err)
	}

	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章创建成功！",
	})
}

// GetArticleInfo 获取文章信息
func GetArticleInfo(c *gin.Context) {
	var a model.Article
	fmt.Println(a)

	// 返回表单
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "test"})
}
