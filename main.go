package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/models"
)

// fixme: why ctx?
// todo: move to some other package
// main is not the best place for controllers
type ctx struct {
	tokens dao.TokensDao
}

func (c *ctx) GetTokens(w http.ResponseWriter, request *http.Request) {
	// todo:
	// it's a common scenario to support pagination with limit & offset
	// also would be great to have total count in the response body
	rawValue := request.URL.Query().Get("limit")
	// fixme:
	// 1. we should not parse/validate user inputs here
	// 2. what if limit is missing?
	limitNum, _ := strconv.ParseUint(rawValue, 0, 32)
	tokens, err := c.tokens.GetTokens(limitNum)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("internal server error"))
	} else {
		// todo:
		// would be great if we had serialize function(or method)
		// in order to hide the exact implementation
		// so it will be possible to change serialization strategies
		data, _ := json.Marshal(tokens)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
	}
}

func (c *ctx) SaveTokens(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	newToken := &models.Token{}

	_ = json.Unmarshal(body, newToken)
	saved, err := c.tokens.SaveToken(newToken)
	if saved.Count == newToken.Count {
		w.WriteHeader(http.StatusCreated)
	} else {
		// todo: change to OK when we return body
		w.WriteHeader(http.StatusNoContent)
	}
	//todo: return the created/updated entity in body
}

func InitRouter(c *ctx) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tokens", c.GetTokens).Methods("GET")
	router.HandleFunc("/tokens", c.SaveTokens).Methods("POST")

	return router
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to init database with error %+v", err)
	}
	router := InitRouter(&ctx{tokens: dao.NewTokensDao(db)})
	// fixme: handle error here (e.g. EADDRINUSE)
	_ = http.ListenAndServe(":8000", router)
}

func initDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=postgres sslmode=disable", "localhost", "5432")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Token{})
	return db, nil
}
