package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryCode	string `gorm:"size:10"`
	Name		string `gorm:"size:50"`
	Note		string
}
