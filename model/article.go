package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	ArticleName    string `gorm:"varchar(20);not null"`
	ArticleContent string `gorm:"varchar(200);not null"`
}
