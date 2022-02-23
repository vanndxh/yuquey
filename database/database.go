package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"yuquey/model"
)

var DB *gorm.DB

func init() {
	var err error
	host := "localhost"
	port := 1101
	database := "postgres"
	username := "postgres"
	password := "pgsql1101"
	psgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	DB, err = gorm.Open("postgres", psgInfo)
	DB.LogMode(true)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	// 自动建表
	DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Timeline{}, &model.SupportCount{}, &model.Like{},
		&model.Collection{}, &model.Feedback{}, &model.Team{}, &model.Comment{}, &model.Follow{}, &model.Message{})
}
