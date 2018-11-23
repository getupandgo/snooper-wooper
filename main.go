package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"github.com/getupandgo/snooper-wooper/mock"
)

func GetTokens (w http.ResponseWriter, request *http.Request) {
	rawValue := request.URL.Query().Get("limit")

	limitNum, _ := strconv.ParseUint(rawValue, 0, 32)
	
	tokens := mock.GetTokens(limitNum)

	data, _ := json.Marshal(tokens)

	w.WriteHeader(200)
	w.Write(data)
}

func SaveTokens (w http.ResponseWriter, request *http.Request) {

}

func InitRouter () (*mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/tokens", GetTokens).Methods("GET")
	router.HandleFunc("/tokens", SaveTokens).Methods("POST")

	return router
}

func main () {
	router := InitRouter()
	
	fmt.Println("Hello Vanya")
	
	http.ListenAndServe(":8000", router)
}
