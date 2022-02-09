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
	articleAuthor, err := strconv.Atoi(c.PostForm("articleAuthor"))
	if err != nil {
		fmt.Println(err)
		return
	}
	isInTrash, err2 := strconv.Atoi(c.PostForm("isInTrash"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	result := database.DB.Find(&a, "article_author=? AND is_in_trash=?", articleAuthor, isInTrash)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "The articles are not exist."})
	}
}

func SearchArticle(c *gin.Context) {
	// 获取数据
	var a []model.Article
	searchValue := c.PostForm("searchValue")
	// 根据searchValue模糊搜索
	result := database.DB.Find(&a, "article_name like ?", "%"+searchValue+"%")
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": a})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "没有找到文章~"})
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
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": "finderr " + result.Error.Error()})
		return
	}
	articleNow := u.ArticleAmount
	database.DB.Model(&u).Update("article_amount", articleNow+1)
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "文章创建成功！",
	})
}

func DeleteArticle(c *gin.Context) {
	// 获取文章id
	var a model.Article
	articleId := c.PostForm("articleId")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(userId)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err})
		return
	}
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
	articleId := c.PostForm("articleId")
	// 查找对应文章
	result := database.DB.Find(&a, "article_id=?", articleId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
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
	articleId := c.PostForm("articleId")
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
	handle := c.PostForm("handle")

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
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "库中没有文章"})
	}
}
