package handler

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	dbconn "github.com/att-innovate/boomer/dbconnector"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user dbconn.User

	err := user.Parse(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = user.Login()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user dbconn.User

	if err := user.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := user.Add(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var user dbconn.User
	//var err error

	user.Id, _ = strconv.ParseInt(vars["Userid"], 10, 64)

	if err := user.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []dbconn.User
	var user dbconn.User

	if err := user.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	users, err := user.Gets()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.ParseInt(vars["Userid"], 10, 64)

	var user dbconn.User

	if err := user.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if uid != user.Id {
		err := fmt.Errorf("Error: Userid mismatch")
		http.Error(w, err.Error(), 500)
		return
	}

	if err := user.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var user dbconn.User
	//var err error

	user.Id, _ = strconv.ParseInt(vars["Userid"], 10, 64)

	if err := user.Delete(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func GetUserCpe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var cpes []dbconn.Cpe
	var cpe dbconn.Cpe
	var err error

	cpe.Userid, err = strconv.ParseInt(vars["Userid"], 10, 64)

	cpes, err = cpe.Gets()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cpes); err != nil {
		panic(err)
	}

}
