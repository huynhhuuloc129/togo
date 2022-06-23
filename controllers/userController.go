package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/huynhhuuloc129/todo/models"
)

func ResponeAllUser(w http.ResponseWriter, r *http.Request) { // Get all user from database
	w.Header().Set("Content-Type", "application/json")
	users, err := models.GetAllUser()
	if err != nil {
		http.Error(w, "get all user failed", http.StatusFailedDependency)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "encode failed", 500)
		return
	}
}

func ResponeOneUser(w http.ResponseWriter, r *http.Request) { // Get one user from database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url
	user, ok := models.CheckIDAndReturn(id)
	if !ok {
		http.Error(w, "id invalid", http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "encode failed", http.StatusFailedDependency)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) { // Create a new user
	var user models.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "decode failed", http.StatusFailedDependency)
		return
	}

	err = models.InsertUser(user)
	if err != nil {
		http.Error(w, "insert user failed", http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "encode failed", http.StatusCreated)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) { // Delete one user from database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url

	_, ok := models.CheckIDAndReturn(id)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	err := models.DeleteAllTaskFromUser(id)
	if err != nil {
		http.Error(w, "delete task of user failed", http.StatusFailedDependency)
		return
	}

	err = models.DeleteUser(id)
	if err != nil {
		http.Error(w, "delete user failed", http.StatusFailedDependency)
		return
	}

	w.Write([]byte("message: delete success"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) { // Update one user already exist in database
	id, _ := strconv.Atoi(fmt.Sprintf("%v", context.Get(r, "id"))) // get id from url

	var newUser models.NewUser
	_, ok := models.CheckIDAndReturn(id)
	if !ok {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}

	err = models.UpdateUser(newUser, id)
	if err != nil {
		http.Error(w, "update user failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, "encode failed.", http.StatusBadRequest)
		return
	}
}
