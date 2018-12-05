package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model `json:"-"`
	Text       string `gorm:"not null;unique" json:"text"`
	Count      uint64 `json:"count"`
}
