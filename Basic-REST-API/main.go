package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Code to fetch and return a list of users
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Code to fetch and return a single user by ID
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
