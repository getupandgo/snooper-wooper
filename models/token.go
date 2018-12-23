package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model `json:"-" binding:"-"`
	Text       string `json:"text"	binding:"required,min=1,max=50"	gorm:"not null;unique"`
	Count      uint64 `json:"count" binding:"required,min=1"`
}
