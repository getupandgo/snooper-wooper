package token

import (
	"bytes"
	"encoding/json"
	"github.com/getupandgo/snooper-wooper/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultTokensRetrieval(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tokens", nil)
	response := httptest.NewRecorder()

	InitRouter().ServeHTTP(response, request)

	if response.Body.Len() != 0 {
		parsedTokens := []mock.Token{}

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

	InitRouter().ServeHTTP(response, request)

	if response.Body.Len() != 0 {
		parsedTokens := []mock.Token{}

		body, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(body, &parsedTokens)

		assert.Equal(t, 11, len(parsedTokens), "Len is expected to be 10 by default")
	} else {
		assert.Fail(t, "Empty responce for 11 elems")
	}
}

func TestTokenCreation(t *testing.T) {
	newToken := mock.Token{"dog", 1}
	encoded, _ := json.Marshal(newToken)

	request, _ := http.NewRequest("POST", "/tokens", bytes.NewReader(encoded))
	response := httptest.NewRecorder()

	InitRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok is expected")
}
