package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		ArticleId:      100001,
		ArticleName:    articleName,
		ArticleContent: articleContent,
	}
	err := database.DB.Create(&newArticle).Error
	if err != nil {
		fmt.Println(err)
	}

	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章创建成功！",
	})
}

func DeleteArticle(c *gin.Context) {
	// 获取文章id
	var a model.Article
	id, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	// 删除操作
	result := database.DB.Delete(&a, "article_id=?", id)
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
	id := c.Param("articleId")

	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", id)
	fmt.Println(result)

	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount
	returnJSON["Hot"] = a.Hot
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}

func GetHotArticle(c *gin.Context) {
	// 获取数据
	var a model.Article
	id := c.Param("articleId")

	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", id)
	fmt.Println(result)

	// 返回表单
	returnJSON := make(map[string]interface{})
	returnJSON["ArticleName"] = a.ArticleName
	returnJSON["ArticleContent"] = a.ArticleContent
	returnJSON["LikeAmount"] = a.LikeAmount
	returnJSON["StarAmount"] = a.StarAmount
	returnJSON["Hot"] = a.Hot
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": returnJSON})
}

func UpdateArticle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "success", "data": "test"})
}
