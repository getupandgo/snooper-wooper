package dao

//go:generate mockgen -destination=../mock/token.go -package=dao github.com/getupandgo/snooper-wooper/dao TokensDao

import (
	"github.com/jinzhu/gorm"

	"github.com/getupandgo/snooper-wooper/models"
)

type (
	TokensDao interface {
		FindToken(text string) (*models.Token, error)
		CreateToken(t *models.Token) (*models.Token, error)
		UpdateToken(t *models.Token) (*models.Token, error)
		GetTopTokens(limit uint64) ([]models.Token, error)
	}

	tokensDao struct {
		db *gorm.DB
	}
)

func (dao tokensDao) CreateToken(t *models.Token) (*models.Token, error) {
	if err := dao.db.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (dao tokensDao) FindToken(text string) (*models.Token, error) {
	found := &models.Token{}
	if err := dao.db.Where("text = ?", text).First(found).Error; err != nil {
		return nil, err
	}
	return found, nil
}

func (dao tokensDao) UpdateToken(t *models.Token) (*models.Token, error) {
	if err := dao.db.Save(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (dao tokensDao) GetTopTokens(limit uint64) ([]models.Token, error) {
	tokens := make([]models.Token, limit)
	if err := dao.db.Order("count DESC").Limit(limit).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func NewTokensDao(db *gorm.DB) TokensDao {
	return &tokensDao{db}
}
