package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/models"
)

type (
	TokenController struct {
		tokens dao.TokensDao
	}
	GetTopTokensQueryString struct {
		Limit uint64 `form:"limit,default=10"`
	}
)

func New(dao dao.TokensDao) TokenController {
	return TokenController{dao}
}

func (ctrl TokenController) GetTopTokens(c *gin.Context) {
	query := &GetTopTokensQueryString{}
	if err := c.BindQuery(query); err != nil {
		c.Error(err)
		return
	}

	if tokens, err := ctrl.tokens.GetTopTokens(query.Limit); err != nil {
		c.Error(err)
		return
	} else {
		c.JSON(http.StatusOK, tokens)
	}
}

func (ctrl TokenController) UpsertToken(c *gin.Context) {
	token := &models.Token{}
	if err := c.ShouldBindJSON(token); err != nil {
		c.Error(err)
		return
	}
	t, err := ctrl.tokens.FindToken(token.Text)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.Error(err)
			return
		}
		t, err = ctrl.tokens.CreateToken(token)
	} else {
		t.Count += token.Count
		t, err = ctrl.tokens.UpdateToken(t)
	}
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, t)
}
