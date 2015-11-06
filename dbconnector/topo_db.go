package dbconnector

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Topo struct {
	Id        int64     `json:"topoid"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	Cpeid     int64     `json:"cpeid"`
	Userid    int64     `json:"userid"`
	Name      string    `json:"name"`
	Topogui   string    `json:"topogui"`
	Topoorch  string    `json:"topoorch"`
}

func (topo *Topo) Parse(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &topo); err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}

func (topo *Topo) Add() error {
	var err error
	err = nil

	if err = db.Create(&topo).Error; err != nil {
		return err
	}
	db.Last(&topo)

	return err
}

func (topo *Topo) Update() error {

	var err error
	err = nil

	if err = db.Table("topos").Where("id = ?", topo.Id).Updates(&topo).Error; err != nil {
		return err
	}
	return err
}

func (topo *Topo) Get() error {
	var err error
	err = nil

	if err = db.Where(&topo).Find(&topo).Error; err != nil {
		return err
	}
	return err
}

func (topo *Topo) Gets() ([]Topo, error) {

	var topos []Topo
	var err error
	err = nil

	if err = db.Where(&topo).Find(&topos).Error; err != nil {
		return topos, err
	}

	return topos, err
}

func (topo *Topo) Delete() error {
	var err error
	err = nil

	if err = db.Where("id = ?", topo.Id).Delete(&topo).Error; err != nil {
		return err
	}
	return err
}
