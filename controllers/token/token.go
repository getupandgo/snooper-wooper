package token

import (
	"net/http"
	"strconv"

	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/models"
	"github.com/gin-gonic/gin"
)

const defaultLimit = "10"

type TokenController struct {
	tokens dao.TokensDao
}

func New(dao dao.TokensDao) TokenController {
	return TokenController{dao}
}

func (ctrl TokenController) GetTopTokens(c *gin.Context) {
	limit, _ := strconv.ParseUint(c.DefaultQuery("limit", defaultLimit), 10, 64)
	tokens, _ := ctrl.tokens.GetTopTokens(limit)
	c.JSON(http.StatusOK, tokens)
}

func (ctrl TokenController) UpsertToken(c *gin.Context) {
	token := &models.Token{}
	if err := c.ShouldBindJSON(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	t, err := ctrl.tokens.SaveToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	c.JSON(http.StatusOK, t)
}
