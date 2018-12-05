package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/mock"
)

type ctx struct {
	TokenDS dao.TokensDAO
}

func (c *ctx) GetTokens(w http.ResponseWriter, request *http.Request) {
	rawValue := request.URL.Query().Get("limit")

	limitNum, _ := strconv.ParseUint(rawValue, 0, 32)

	tokens := mock.GetTokens(limitNum)

	data, _ := json.Marshal(tokens)

	w.WriteHeader(200)
	w.Write(data)
}

func (c *ctx) SaveTokens(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	newToken := mock.Token{}
	_ = json.Unmarshal(body, &newToken)

	mock.SaveToken(newToken)
	w.WriteHeader(200)
}

func InitRouter(c *ctx) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tokens", c.GetTokens).Methods("GET")
	router.HandleFunc("/tokens", c.SaveTokens).Methods("POST")

	return router
}

func main() {
	// creaet context struct and connect methods to it
	context := ctx{}
	token_dao, _ := dao.Connect()

	context

	router := InitRouter()

	http.ListenAndServe(":8000", router)
}
