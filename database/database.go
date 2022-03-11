package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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

}
