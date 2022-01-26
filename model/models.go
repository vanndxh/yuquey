package model

type User struct {
	UserId        int    `gorm:"primary_key;not null;unique"`
	Username      string `gorm:"varchar(20);not null"`
	Password      string `gorm:"size:255;not null"`
	UserInfo      string `gorm:"varchar(200)"`
	ArticleAmount int    `gorm:"int8;not null"`
	LikeTotal     int8   `gorm:"int8;not null"`
}

type Article struct {
	ArticleId      int    `gorm:"primary_key;not null"`
	ArticleName    string `gorm:"varchar(20);not null"`
	ArticleContent string `gorm:"varchar(200);not null"`
	LikeAmount     int    `gorm:"not null"`
	StarAmount     int    `gorm:"not null"`
	Hot            int    `gorm:"not null"`
}

type Timeline struct {
	Title   string `gorm:"varchar(20);not null"`
	Content string `gorm:"varchar(100);not null"`
	Time    string `gorm:"varchar(20);not null"`
	Type    string `gorm:"varchar(20)"`
}

type SupportCount struct {
	Count int `gorm:"not null"`
}

type Like struct {
	UserId    int `gorm:"not null"`
	ArticleId int `gorm:"not null"`
}

type Star struct {
	UserId    int `gorm:"not null"`
	ArticleId int `gorm:"not null"`
}
