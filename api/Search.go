package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yuquey/database"
	"yuquey/model"
	"yuquey/util"
)

func Search(c *gin.Context) { // 根据搜索内容模糊查询
	searchValue := c.DefaultQuery("searchValue", "")
	handle := c.DefaultQuery("handle", "0") // 0-article 1-user 2-team

	if handle == "0" {
		// 计算最新hot
		util.CalculateHot()
		// 模糊搜索文章名
		var as []model.Article
		result := database.DB.Where("secret=?", 0).Order("hot desc").Find(&as, "article_name like ?", "%"+searchValue+"%")
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
	} else if handle == "1" {
		var us []model.User
		res := database.DB.Order("user_id").Find(&us, "username like ?", "%"+searchValue+"%")
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": us})
	} else if handle == "2" {
		var ts []model.Team
		res := database.DB.Order("team_id").Find(&ts, "team_name like ?", "%"+searchValue+"%")
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
			return
		}
		for i := range ts {
			var u model.User
			database.DB.Find(&u, "user_id=?", ts[i].TeamLeader)
			ts[i].LeaderName = u.Username
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ts})
	}

}
