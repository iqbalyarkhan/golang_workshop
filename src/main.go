package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//WelcomeMessage to hold message for welcome page
type WelcomeMessage struct {
	Message string `json:"message"`
}

//Users struct to hold user information
type Users struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

var globalUsers []Users

//GetUsers function to return all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(globalUsers)
}

//GetUser function to get user with specific ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range globalUsers {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&Users{})
}

//CreateUser function to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	globalUsers = append(globalUsers, user)
	json.NewEncoder(w).Encode(globalUsers)
}

//DeleteUser function to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range globalUsers {
		if user.ID == params["id"] {
			globalUsers = append(globalUsers[:index], globalUsers[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(globalUsers)
	}
}

//Welcome function display welcome message
func Welcome(w http.ResponseWriter, r *http.Request) {
	msg := "We're live"
	json.NewEncoder(w).Encode(msg)
}

func main() {
	router := mux.NewRouter()
	globalUsers = append(globalUsers, Users{"1", "Ryan", "Reynolds", 35})
	globalUsers = append(globalUsers, Users{"2", "Ray", "Donovan", 30})
	globalUsers = append(globalUsers, Users{"3", "Tom", "Brady", 41})

	router.HandleFunc("/", Welcome).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/createUser", CreateUser).Methods("POST")
	router.HandleFunc("/deleteUser/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
