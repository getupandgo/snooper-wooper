package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	Text  string `gorm:"not null;unique"`
	Count uint64
}
