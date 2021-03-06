package model

import (
	"time"
)

type User struct {
	UserId         int       `gorm:"primary_key;AUTO_INCREASE"`
	Username       string    `gorm:"varchar(10);not null"`
	Password       string    `gorm:"not null"`
	UserInfo       string    `gorm:"varchar(60)"`
	Sex            int       `gorm:"not null"`
	Vip            time.Time `gorm:"not null"`
	Authentication string    `gorm:""`
	ArticleAmount  int       `gorm:"not null"`
	LikeAmount     int       `gorm:"not null"`
	FollowAmount   int       `gorm:"not null"`
	FollowerAmount int       `gorm:"not null"`
	ReadNotice     int       `gorm:"not null"`
	Violation      int       `gorm:"not null"`
	LatestTaskTime time.Time `gorm:"not null"`
}

type Article struct {
	ArticleId        int       `gorm:"primary_key;AUTO_INCREASE"`
	ArticleName      string    `gorm:"varchar(15);not null"`
	ArticleContent   string    `gorm:"varchar(2000)"`
	ArticleAuthor    int       `gorm:"not null"`
	AuthorName       string    `gorm:"-"`
	Secret           int       `gorm:"not null"`
	Time             time.Time `gorm:"not null"`
	ViewAmount       int       `gorm:"not null"`
	LikeAmount       int       `gorm:"not null"`
	CollectionAmount int       `gorm:"not null"`
	IsInTrash        int       `gorm:"not null"` // 0-not 1-yes
	Hot              float64   `gorm:"not null"`
	Tag              string    `gorm:""`
}

type Team struct {
	TeamId     int    `gorm:"primary_key;AUTO_INCREASE"`
	TeamName   string `gorm:"varchar(10);not null"`
	TeamNotice string `gorm:"varchar(40)"`
	TeamLeader int    `gorm:"not null"`
	LeaderName string `gorm:"-"`
}

type TeamUser struct {
	UserId        int       `gorm:"not null"`
	Username      string    `gorm:"-"`
	TeamId        int       `gorm:"not null"`
	TeamName      string    `gorm:"-"`
	Position      int       `gorm:"not null"` // 0-leader 1-manager 2-member
	PositionName  string    `gorm:"-"`
	Punch         int       `gorm:"not null"`
	LastPunchTime time.Time `gorm:"not null"`
}

type Comment struct {
	CommentId      int       `gorm:"primary_key;AUTO_INCREASE"`
	UserId         int       `gorm:"not null"`
	Username       string    `gorm:"-"`
	UserVip        time.Time `gorm:"-"`
	UserAuth       string    `gorm:"-"`
	ArticleId      int       `gorm:"not null"`
	ArticleName    string    `gorm:"-"`
	CommentContent string    `gorm:"not null"`
}

type Message struct {
	MessageId int       `gorm:"primary_key;AUTO_INCREASE"`
	UserId    int       `gorm:"not null"`
	Type      int       `gorm:"not null"` // 0-all 1-like 2-collection 3-comment 4-follow 5-other
	Op        int       `gorm:"not null"`
	ArticleId int       `gorm:""`
	Content   string    `gorm:""`
	Read      int       `gorm:"not null"` // 1-no 2-yes
	Time      time.Time `gorm:"not null"`
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

type Notice struct {
	NoticeContent string    `gorm:"not null"`
	Time          time.Time `gorm:""`
}

type Feedback struct {
	FeedbackId   int    `gorm:"primary_key;AUTO_INCREASE"`
	FeedbackInfo string `gorm:"varchar(200);not null"`
	UserId       int    `gorm:"not null"`
}
