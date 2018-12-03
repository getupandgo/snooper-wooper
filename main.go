package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"
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

func InitRouter () (*mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/tokens", GetTokens).Methods("GET")
	router.HandleFunc("/tokens", SaveTokens).Methods("POST")

	return router
}

func main () {
	router := InitRouter()
	
	http.ListenAndServe(":8000", router)
}
