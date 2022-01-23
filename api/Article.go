package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

	//创建文章
	newArticle := model.Article{
		ArticleId:      "100000001",
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

func DeleteArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	result := database.DB.Delete(&model.Article{}, articleId)
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
	var a model.Article
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}

func GetHotArticle(c *gin.Context) {
	var a model.Article
	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}
