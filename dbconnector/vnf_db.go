package dbconnector

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Vnf struct {
	Id        int64     `json:"vnfid"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	Name      string    `json:"name"`
	Vnftype   string    `json:"vnftype"` // Router, Firewall, Cache, etc
	Flavor    string    `json:"flavor"`  // docker, LXC, VM, etc
	Image     string    `json:"image"`
	Version   string    `json:"version"`
	Template  string    `json:"template"`
	Distro    string    `json:"distro"`
	Release   string    `json:"release"`
	Arch      string    `json:"arch"`
	Startcmd  string    `json:"startcmd"`
}

func (vnf *Vnf) Parse(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &vnf); err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}

func (vnf *Vnf) Add() error {
	var err error

	err = nil
	if err = db.Create(&vnf).Error; err != nil {
		return err
	}
	db.Last(&vnf)

	return err
}

func (vnf *Vnf) Update() error {
	var err error

	err = nil
	if err = db.Table("vnfs").Where("id = ?", vnf.Id).Updates(&vnf).Error; err != nil {
		return err
	}
	return err
}

func (vnf *Vnf) Get() error {
	var err error

	err = nil
	if err = db.Where(&vnf).Find(&vnf).Error; err != nil {
		return err
	}
	return err
}

func (vnf *Vnf) Gets() ([]Vnf, error) {

	var vnfs []Vnf
	var err error

	err = nil
	if err = db.Where(&vnf).Find(&vnfs).Error; err != nil {
		return vnfs, err
	}

	return vnfs, err
}

func (vnf *Vnf) Delete() error {
	var err error

	err = nil
	if err = db.Where("id = ?", vnf.Id).Delete(&vnf).Error; err != nil {
		return err
	}
	return err
}
