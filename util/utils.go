package util

import (
	"time"
	"yuquey/database"
	"yuquey/model"
)

func CalculateHot() {
	var as []model.Article
	database.DB.Find(&as, "is_in_trash=? AND secret=?", 0, 0)
	for i := range as {
		// 获取文章评论数
		var cs []model.Comment
		res := database.DB.Find(&cs, "article_id=?", as[i].ArticleId)
		// 获取已经过了几天
		days := time.Now().Sub(as[i].Time).Hours()/24 + 1
		// （浏览量，点赞，评论，收藏）加权求和
		score := float64(as[i].ViewAmount + as[i].LikeAmount*10 + int(res.RowsAffected)*20 + as[i].CollectionAmount*30)
		// 计算作者额外权重
		var u model.User
		database.DB.Find(&u, "user_id=?", as[i].ArticleAuthor)
		var author = 1.0
		author = float64(1 + (u.ArticleAmount*5+u.LikeAmount+u.FollowAmount*2+u.FollowerAmount*10)/100)
		if u.Vip.After(time.Now()) {
			author *= 1.5 // VIP用户权重再1.5倍
		}
		if len(u.Authentication) != 0 {
			author *= 1.5 // 如果是认证号，再再1.5
		}
		// 汇总计算
		newHot := 1.0 * score * author / days
		// 更新
		var a model.Article
		if newHot != as[i].Hot {
			database.DB.Model(&a).Where("article_id=?", as[i].ArticleId).Update("hot", newHot)
		}
	}
}
