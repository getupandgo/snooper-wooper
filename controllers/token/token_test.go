package token_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/getupandgo/snooper-wooper/controllers"
	dao "github.com/getupandgo/snooper-wooper/mock"
	"github.com/getupandgo/snooper-wooper/models"
)

func TestDefaultTokensRetrieval(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tokens", nil)
	response := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	x := dao.NewMockTokensDao(ctrl)

	x.EXPECT().GetTopTokens(uint64(10)).Return(dao.SampleTokens[:10], nil).Times(1)

	controllers.InitRouter(x).ServeHTTP(response, request)

	if response.Body.Len() != 0 {
		parsedTokens := []models.Token{}

		body, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(body, &parsedTokens)

		assert.Equal(t, 10, len(parsedTokens), "Len is expected to be 10 by default")
	} else {
		assert.Fail(t, "Empty responce for default ")
	}
}

func TestLimitTokensRetrieval(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tokens?limit=11", nil)
	response := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	x := dao.NewMockTokensDao(ctrl)

	x.EXPECT().GetTopTokens(uint64(11)).Return(dao.SampleTokens[:11], nil)

	controllers.InitRouter(x).ServeHTTP(response, request)

	if response.Body.Len() != 0 {
		parsedTokens := []models.Token{}

		body, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(body, &parsedTokens)

		assert.Equal(t, 11, len(parsedTokens), "Len is expected to be 10 by default")
	} else {
		assert.Fail(t, "Empty responce for 11 elems")
	}
}

func TestTokenCreation(t *testing.T) {
	newToken := models.Token{Text: "dog", Count: 1}
	encoded, _ := json.Marshal(newToken)

	request, _ := http.NewRequest("POST", "/tokens", bytes.NewReader(encoded))
	response := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daoMock := dao.NewMockTokensDao(ctrl)
	daoMock.EXPECT().FindToken(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
	daoMock.EXPECT().CreateToken(gomock.Any()).Return(&newToken, nil)

	controllers.InitRouter(daoMock).ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Ok is expected")
}
