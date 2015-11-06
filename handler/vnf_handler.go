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

func AddVnf(w http.ResponseWriter, r *http.Request) {

	var vnf dbconn.Vnf

	err := vnf.Parse(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = vnf.Add(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vnf); err != nil {
		panic(err)
	}

}

func GetVnf(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var vnf dbconn.Vnf
	var err error

	vnf.Id, err = strconv.ParseInt(vars["Vnfid"], 10, 64)

	if err = vnf.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vnf); err != nil {
		panic(err)
	}

}

func GetVnfs(w http.ResponseWriter, r *http.Request) {

	var vnfs []dbconn.Vnf
	var vnf dbconn.Vnf

	vnfs, err := vnf.Gets()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vnfs); err != nil {
		panic(err)
	}

}

func UpdateVnf(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	vid, _ := strconv.ParseInt(vars["Vnfid"], 10, 64)

	var vnf dbconn.Vnf

	if err := vnf.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if vid != vnf.Id {
		err := fmt.Errorf("Error: Vnfid mismatch")
		http.Error(w, err.Error(), 500)
		return
	}

	if err := vnf.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vnf); err != nil {
		panic(err)
	}

}

func DeleteVnf(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var vnf dbconn.Vnf
	var err error

	vnf.Id, err = strconv.ParseInt(vars["Vnfid"], 10, 64)

	if err = vnf.Delete(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
