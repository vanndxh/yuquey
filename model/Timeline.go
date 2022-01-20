package model

import "github.com/jinzhu/gorm"

type Timeline struct {
	gorm.Model
	Title   string `gorm:"varchar(20);not null"`
	Content string `gorm:"varchar(100);not null"`
	Time    string `gorm:"varchar(20);not null"`
	Type    string `gorm:"varchar(20)"`
}
