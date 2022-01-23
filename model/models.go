package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId        string `gorm:"varchar(6);not null;unique"`
	Username      string `gorm:"varchar(20);not null"`
	Password      string `gorm:"size:255;not null"`
	UserInfo      string `gorm:"varchar(200)"`
	ArticleAmount int8   `gorm:"int8;not null"`
	LikeTotal     int8   `gorm:"int8;not null"`
}

type Article struct {
	gorm.Model
	ArticleId      string `gorm:"varchar(9);not null;unique"`
	ArticleName    string `gorm:"varchar(20);not null"`
	ArticleContent string `gorm:"varchar(200);not null"`
	LikeAmount     int8   `gorm:"not null"`
	StarAmount     int8   `gorm:"not null"`
}

type Timeline struct {
	gorm.Model
	Title   string `gorm:"varchar(20);not null"`
	Content string `gorm:"varchar(100);not null"`
	Time    string `gorm:"varchar(20);not null"`
	Type    string `gorm:"varchar(20)"`
}

type SupportCount struct {
	gorm.Model
	Count int8 `gorm:"int8;not null"`
}
