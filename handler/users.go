package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"requestmanagment/api"
	"requestmanagment/models"
)

type UsersHandler struct {
}

var Users = map[string]models.User{}

func (u *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method is  : %s \n", r.Method)
	switch {
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		GetUser(w, r)
		return
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0:
		GetUsers(w, r)
		return
	case r.Method == http.MethodPost:
		CreateUser(w, r)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ApiKey := r.Header.Get("X-API-KEY")

	for k, v := range r.Header {
		fmt.Println(k, v)
	}

	if ApiKey != "234242422423" {
		w.WriteHeader(http.StatusUnauthorized)

		// Return Data
		api.SetResult(http.StatusUnauthorized, nil, nil, w)
		// _, err := fmt.Fprintf(w, "Invalid Api Key")
		// if err != nil {
		// 	return
		// }
		return
	}

	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		api.SetResult(http.StatusInternalServerError, nil, err, w)

		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Println(json.NewDecoder(r.Body))
		// fmt.Fprintf(w, "Decode Error : %v", err)
		return
	}

	//add to database
	if _, exists := Users[strconv.Itoa(user.Id)]; exists {
		api.SetResult(http.StatusConflict, nil, errors.New("user exists"), w)
		return
	}
	Users[strconv.Itoa(user.Id)] = *user

	api.SetResult(http.StatusOK, *user, nil, w)
	// w.WriteHeader(200)
	// fmt.Fprintf(w, "User Created \n %v , %s , %v", user.Id, user.Fullname, user.Age)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	user, exists := Users[id]

	if exists {
		api.SetResult(http.StatusOK, user, nil, w)
	} else {
		api.SetResult(http.StatusNotFound, nil, nil, w)
	}
	// json.NewEncoder(w).Encode(user)

	// _, err := fmt.Fprintf(w, "User Is : %s", id)
	// if err != nil {
	// 	return
	// }
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	api.SetResult(http.StatusOK, Users, nil, w)

	// _, err := fmt.Fprintf(w, "User all users")
	// if err != nil {
	// 	return
	// }
}
