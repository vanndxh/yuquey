package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"time"
	"yuquey/model"
)

// 生产环境
// 5432
//const postgresHost = "xhw_postgres"

// 本地开发
// 1101
const postgresHost = "localhost"

var DB *gorm.DB

func init() {
	// postgresDB连接
	var err error
	psgInfo := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
		postgresHost, "postgres", "1101", "yuquey1101", "postgres")
	DB, err = gorm.Open("postgres", psgInfo)
	DB.LogMode(true)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	// 自动建表
	DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Timeline{}, &model.SupportCount{}, &model.Like{},
		&model.Collection{}, &model.Feedback{}, &model.Team{}, &model.Comment{}, &model.Follow{}, &model.Message{},
		&model.TeamUser{}, &model.Notice{})

	// 初始化sc
	var sc model.SupportCount
	res := DB.Find(&sc)
	if res.RowsAffected == 0 {
		newSupportCount := model.SupportCount{
			Count: 1,
		}
		err := DB.Create(&newSupportCount).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 初始化notice
	var n model.Notice
	res2 := DB.Find(&n)
	if res2.RowsAffected == 0 {
		newNotice := model.Notice{
			NoticeContent: "暂无新公告~",
		}
		err := DB.Create(&newNotice).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 初始化管理员账号
	adminUser := model.User{
		Username:       "Vanndxh",
		Password:       "123456",
		UserInfo:       "暂无~",
		Authentication: "网站创始人",
	}
	err2 := DB.Create(&adminUser).Error
	if err2 != nil {
		log.Println(err2)
	}

	// 创建第一篇文章留用
	firstArticle := model.Article{
		ArticleName:    "小黑屋使用攻略",
		ArticleContent: "等创完了再来写~",
		ArticleAuthor:  1,
		Time:           time.Now(),
	}
	DB.Create(&firstArticle)

	// 初始化文章库
	var i = 0
	for i = 0; i < 30; i++ {
		newArticle := model.Article{
			ArticleName:    "初始化文章" + strconv.Itoa(i),
			ArticleContent: "本文章为初始化需要，没有实际内容，没有作者！",
			ArticleAuthor:  99999,
		}
		DB.Create(&newArticle)
	}

}
