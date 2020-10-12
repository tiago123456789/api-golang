package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type People struct {
	Name string `json:name`
	Age int `json:age`
}

var peoples []People

func sendSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	successResponse := SuccessResponse{ Message: "Success!!!!!" }
	json.NewEncoder(w).Encode(&successResponse)
}

func createPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var people People
	_ = json.NewDecoder(r.Body).Decode(&people)
	peoples = append(peoples, people);
	json.NewEncoder(w).Encode(&peoples);
}

func getPeoples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&peoples);
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", sendSuccess).Methods("GET")
	router.HandleFunc("/peoples", createPeople).Methods("POST")
	router.HandleFunc("/peoples", getPeoples).Methods("GET")


	http.ListenAndServe(":8000", router);
}