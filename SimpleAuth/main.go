package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users []User
var mutex sync.Mutex

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	for _, user := range users {
		if user.FirstName == newUser.FirstName && user.LastName == newUser.LastName {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User added successfully")
	fmt.Print(users)
}

func fetchingUserByName(users []User, targetUser User) (User, error) {
	for _, user := range users {
		if user.FirstName == targetUser.FirstName && user.LastName == targetUser.LastName {
			fmt.Printf("Found user: %s %s\n", user.FirstName, user.LastName)
			return user, nil
		}
	}
	fmt.Printf("User %s %s not found\n", targetUser.FirstName, targetUser.LastName)
	return User{}, errors.New("user not found")
}

func gettingUserByName(w http.ResponseWriter, r *http.Request) {
	var targetUser User
	json.NewDecoder(r.Body).Decode(&targetUser)
	user, err := fetchingUserByName(users, targetUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Error encoding user", http.StatusInternalServerError)
		return
	}
}

func removeUserByName(users []User, targetUser User) ([]User, error) {
	for index, user := range users {
		if user.FirstName == targetUser.FirstName && user.LastName == targetUser.LastName {
			fmt.Printf("Found user: %s %s at index: %d\n", user.FirstName, user.LastName, index)
			users = append(users[:index], users[index+1:]...)
			return users, nil
		}
	}
	fmt.Printf("User %s %s not found\n", targetUser.FirstName, targetUser.LastName)
	return users, errors.New("user not found")

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var targetUser User
	err := json.NewDecoder(r.Body).Decode(&targetUser)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	var errMsg string
	users, err = removeUserByName(users, targetUser)
	if err != nil {
		errMsg = err.Error()
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "User deleted successfully")
}

// function to extract query params
func extractQueryParams(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	fmt.Print(queryString.Has("id"))
	fmt.Fprintf(w, "Query String: %s\n", queryString)
}

func main() {

	mux := http.NewServeMux()
	//making the routes
	mux.HandleFunc("/adddUser", addUser)
	mux.HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if len(users) == 0 {
			http.Error(w, "No users found", http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "Error encoding users", http.StatusInternalServerError)
			return
		}
	})
	mux.HandleFunc("/getUserByName", gettingUserByName)
	mux.HandleFunc("/deleteUser", deleteUser)
	mux.HandleFunc("/query", extractQueryParams)
	//making a server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Simulating fetching users from an API

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
