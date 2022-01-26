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
	database := "yuquey"
	username := "postgres"
	password := "vanndxh1101"
	psgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	DB, err = gorm.Open("postgres", psgInfo)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Timeline{}, &model.SupportCount{}, &model.Like{},
		&model.Star{})
}
