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

func GetCpe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var cpe dbconn.Cpe
	var err error

	cpe.Id, err = strconv.ParseInt(vars["Cpeid"], 10, 64)

	if err = cpe.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cpe); err != nil {
		panic(err)
	}

}

func GetCpes(w http.ResponseWriter, r *http.Request) {

	var cpes []dbconn.Cpe
	var cpe dbconn.Cpe

	if err := cpe.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cpes, err := cpe.Gets()
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

func AddCpe(w http.ResponseWriter, r *http.Request) {

	var cpe dbconn.Cpe

	if err := cpe.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := cpe.Add(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cpe); err != nil {
		panic(err)
	}
}

func UpdateCpe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, _ := strconv.ParseInt(vars["Cpeid"], 10, 64)

	var cpe dbconn.Cpe

	if err := cpe.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if cid != cpe.Id {
		err := fmt.Errorf("Cpeid mismatch")
		http.Error(w, err.Error(), 500)
		return
	}

	if err := cpe.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cpe); err != nil {
		panic(err)
	}
}

func DeleteCpe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var cpe dbconn.Cpe
	var err error

	cpe.Id, err = strconv.ParseInt(vars["Cpeid"], 10, 64)
	err = cpe.Delete()
	if err != nil {
		return
	}
}

func GetCpeTopo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var cpe dbconn.Cpe
	var topo dbconn.Topo
	var err error

	cpe.Id, err = strconv.ParseInt(vars["Cpeid"], 10, 64)

	if err = cpe.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	topo.Id = cpe.Topoid
	if err = topo.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(topo); err != nil {
		panic(err)
	}

}
