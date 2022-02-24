package model

import "time"

type User struct {
	UserId         int    `gorm:"primary_key;AUTO_INCREASE"`
	Username       string `gorm:"varchar(20);not null"`
	Password       string `gorm:"size:255;not null"`
	UserInfo       string `gorm:"varchar(200)"`
	ArticleAmount  int    `gorm:"not null"`
	LikeTotal      int    `gorm:"not null"`
	FollowAmount   int    `gorm:"not null"`
	FollowerAmount int    `gorm:"not null"`
}

type Article struct {
	ArticleId        int    `gorm:"primary_key;AUTO_INCREASE"`
	ArticleName      string `gorm:"varchar(20);not null"`
	ArticleContent   string `gorm:"varchar(200)"`
	ArticleAuthor    int    `gorm:"not null"`
	LikeAmount       int    `gorm:"not null"`
	CollectionAmount int    `gorm:"not null"`
	IsInTrash        int    `gorm:"not null"`
	Hot              int    `gorm:"not null"`
}

type Team struct {
	TeamId       int    `gorm:"primary_key;AUTO_INCREASE"`
	TeamName     string `gorm:"varchar(10);not null"`
	TeamNotice   string `gorm:"varchar(40)"`
	TeamLeader   int    `gorm:"not null"`
	LeaderCount  int    `gorm:"not null"`
	TeamMember1  int    `gorm:""`
	Member1Count int    `gorm:"not null"`
	TeamMember2  int    `gorm:""`
	Member2Count int    `gorm:"not null"`
	TeamMember3  int    `gorm:""`
	Member3Count int    `gorm:"not null"`
	TeamMember4  int    `gorm:""`
	Member4Count int    `gorm:"not null"`
}

type Comment struct {
	CommentId      int    `gorm:"primary_key;AUTO_INCREASE"`
	UserId         int    `gorm:"not null"`
	Username       string `gorm:"-"`
	ArticleId      int    `gorm:"not null"`
	CommentContent string `gorm:"not null"`
}

type Message struct {
	MessageId   int       `gorm:"primary_key;AUTO_INCREASE"`
	UserId      int       `gorm:"not null"`
	Username    string    `gorm:"-"`
	Type        int       `gorm:"not null"`
	TypeName    string    `gorm:"-"`
	Op          int       `gorm:"not null"`
	OpName      string    `gorm:"-"`
	ArticleId   int       `gorm:""`
	ArticleName string    `gorm:"-"`
	Read        int       `gorm:"not null"`
	Time        time.Time `gorm:"not null"`
}

type Timeline struct {
	Title   string `gorm:"varchar(20);not null"`
	Content string `gorm:"varchar(100);not null"`
	Time    string `gorm:"varchar(20);not null"`
	Type    string `gorm:"varchar(20)"`
}

type Follow struct {
	UpId         int    `gorm:"not null"`
	UpName       string `gorm:"-"`
	FollowerId   int    `gorm:"not null"`
	FollowerName string `gorm:"-"`
}

type Like struct {
	UserId    int `gorm:"not null"`
	ArticleId int `gorm:"not null"`
}

type Collection struct {
	UserId    int `gorm:"not null"`
	ArticleId int `gorm:"not null"`
}

type SupportCount struct {
	Count int `gorm:"not null"`
}

type Feedback struct {
	FeedbackInfo string `gorm:"varchar(200);not null"`
}
