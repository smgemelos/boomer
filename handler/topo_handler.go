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

func AddTopo(w http.ResponseWriter, r *http.Request) {

	var topo dbconn.Topo
	var cpe dbconn.Cpe

	if err := topo.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := topo.Add(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cpe.Id = topo.Cpeid
	if err := cpe.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cpe.Topoid = topo.Id
	if err := cpe.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(topo); err != nil {
		panic(err)
	}

}

func GetTopo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var topo dbconn.Topo
	var err error

	topo.Id, err = strconv.ParseInt(vars["Topoid"], 10, 64)

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

func GetTopos(w http.ResponseWriter, r *http.Request) {

	var topos []dbconn.Topo
	var topo dbconn.Topo

	topos, err := topo.Gets()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(topos); err != nil {
		panic(err)
	}

}

func UpdateTopo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tid, _ := strconv.ParseInt(vars["Topoid"], 10, 64)

	var topo dbconn.Topo

	if err := topo.Parse(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if tid != topo.Id {
		err := fmt.Errorf("Topoid mismatch")
		http.Error(w, err.Error(), 500)
		return
	}

	if err := topo.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(topo); err != nil {
		panic(err)
	}

}

func DeleteTopo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var topo dbconn.Topo
	var cpe dbconn.Cpe
	var err error

	topo.Id, err = strconv.ParseInt(vars["Topoid"], 10, 64)

	// Need to clear the TopoID on the associated CPE - set it to Topoid = 0
	if err = topo.Get(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cpe.Id = topo.Cpeid
	cpe.Topoid = 0
	if err = cpe.Update(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Delete the topo from Topos Table
	if err = topo.Delete(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
