package dbconnector

import (
	"encoding/json"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Cpe struct {
	Id        int64     `json:"cpeid"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	New       string    `json:"new"`
	Userid    int64     `json:"userid"`
	Topoid    int64     `json:"topoid"`
	Macaddr   string    `json:"macaddr"`
	Mgmtip    string    `json:"mgmtip"`
	Mgmtgw    string    `json:"mgmtgw"`
	Publicip  string    `json:"publicip"`
	Publicgw  string    `json:"publicgw"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
}

func (cpe *Cpe) Parse(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &cpe); err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}

func (cpe *Cpe) Add() error {

	var err error
	err = nil

	if err = db.Create(&cpe).Error; err != nil {
		return err
	}
	db.Last(&cpe)

	return err
}

func (cpe *Cpe) Update() error {

	var err error
	err = nil

	if err = db.Table("cpes").Where("id = ?", cpe.Id).Updates(&cpe).Error; err != nil {
		return err
	}
	if err = db.Where("id = ?", cpe.Id).Find(&cpe).Error; err != nil {
		return err
	}
	return err
}

func (cpe *Cpe) Get() error {

	var err error
	err = nil

	if err = db.Where(&cpe).Find(&cpe).Error; err != nil {
		return err
	}
	return err
}

func (cpe *Cpe) Gets() ([]Cpe, error) {

	var cpes []Cpe
	var err error
	err = nil

	if err = db.Where(&cpe).Find(&cpes).Error; err != nil {
		return cpes, err
	}

	return cpes, err

}

func (cpe *Cpe) Delete() error {

	var err error
	err = nil

	if err = db.Where("id = ?", cpe.Id).Delete(&cpe).Error; err != nil {
		return err
	}
	return err
}
