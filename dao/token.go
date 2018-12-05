package dao

//go:generate mockgen -destination=../mock/token.go -package=dao github.com/getupandgo/snooper-wooper/dao TokensDao
import (
	"github.com/jinzhu/gorm"

	"github.com/getupandgo/snooper-wooper/models"
)

type (
	TokensDao interface {
		SaveToken(t *models.Token) (*models.Token, error)
		GetTokens(limit uint64) ([]models.Token, error)
	}

	tokensDao struct {
		db *gorm.DB
	}
)

func (dao tokensDao) SaveToken(t *models.Token) (*models.Token, error) {
	tmp := &models.Token{}
	if err := dao.db.FirstOrCreate(tmp, models.Token{Text: t.Text}).Error; err != nil {
		return nil, err
	}
	// fixme: this is a wrong place for "business logic"
	tmp.Count += t.Count
	if err := dao.db.Save(tmp).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (dao tokensDao) GetTokens(limit uint64) ([]models.Token, error) {
	tokens := make([]models.Token, limit)
	if err := dao.db.Limit(limit).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func NewTokensDao(db *gorm.DB) TokensDao {
	return &tokensDao{db}
}
