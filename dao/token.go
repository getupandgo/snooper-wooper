package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"

	"github.com/getupandgo/snooper-wooper/models"
)

type TokensDAO struct {
	db *gorm.DB
}

func (dao TokensDAO) SaveToken(t *models.Token) (*models.Token, error) {
	tmp := &models.Token{}
	if err := dao.db.FirstOrCreate(tmp, models.Token{Text: t.Text}).Error; err != nil {
		return nil, err
	}
	tmp.Count += t.Count
	if err := dao.db.Save(tmp).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (dao TokensDAO) GetTokens(limit uint64) ([]models.Token, error) {
	tokens := make([]models.Token, limit)
	if err := dao.db.Limit(limit).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func Connect() (TokensDAO, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=postgres sslmode=disable", "localhost", "5432")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("error %+v", err)
		os.Exit(1)
	}
	db.LogMode(true)
	dao := TokensDAO{db}
	db.AutoMigrate(&models.Token{})

	return dao, err
}
