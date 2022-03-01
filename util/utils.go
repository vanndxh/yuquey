package util

import (
	"time"
	"yuquey/database"
	"yuquey/model"
)

func CalculateHot() {
	var as []model.Article
	database.DB.Find(&as)
	for i := range as {
		var cs []model.Comment
		res := database.DB.Find(&cs, "article_id=?", as[i].ArticleId)
		days := int(time.Now().Sub(as[i].Time).Hours()/24 + 1)
		score := as[i].ViewAmount + as[i].LikeAmount*10 + int(res.RowsAffected)*20 + as[i].CollectionAmount*30
		newHot := score / days
		var a model.Article
		database.DB.Model(&a).Where("article_id=?", as[i].ArticleId).Update("hot", newHot)
	}
}
