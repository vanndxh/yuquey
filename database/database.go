package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"yuquey/model"
)

// 生产环境
const (
	postgresHost = "xhw_postgres:1101"
)

var DB *gorm.DB

func init() {
	// postgresDB
	var err error
	psgInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, "postgres", "pgsql1101", "postgres")
	DB, err = gorm.Open("postgres", psgInfo)
	DB.LogMode(true)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// 自动建表
	DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Timeline{}, &model.SupportCount{}, &model.Like{},
		&model.Collection{}, &model.Feedback{}, &model.Team{}, &model.Comment{}, &model.Follow{}, &model.Message{},
		&model.TeamUser{})
}
