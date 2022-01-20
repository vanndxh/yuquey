package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId        string `gorm:"varchar(6);not null"`
	Username      string `gorm:"varchar(20);not null"`
	Password      string `gorm:"size:255;not null"`
	UserInfo      string `gorm:"varchar(200)"`
	ArticleAmount int8   `gorm:"int8;not null"`
	LikeTotal     int8   `gorm:"int8;not null"`
}
