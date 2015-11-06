package dbconnector

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Id        int64     `json:"userid"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	Level     int       `json:"level"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

func (user *User) Parse(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}

func (user *User) Login() error {

	var ruser User
	var err error

	err = nil
	if err = db.Where("username = ?", user.Username).Find(&ruser).Error; err != nil {
		if err.Error() == "record not found" {
			err = fmt.Errorf("Error: Bad Username and/or Password.")
		}
		return err
	}

	if ruser.Password == user.Password {
		err = nil
		*user = ruser
	} else {
		err = fmt.Errorf("Error: Bad Username and/or Password.")
	}

	return err
}

func (user *User) Add() error {
	var err error

	err = nil
	if err = db.Create(&user).Error; err != nil {
		return err
	}
	db.Last(&user)

	return err

}

func (user *User) Update() error {

	var err error

	err = nil
	if err = db.Table("users").Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return err
	}
	return err
}

func (user *User) Get() error {

	var err error

	err = nil
	if err = db.Where(&user).Find(&user).Error; err != nil {
		return err
	}
	return err
}

func (user *User) Gets() ([]User, error) {

	var users []User
	var err error

	err = nil
	if err = db.Where(&user).Find(&users).Error; err != nil {
		return users, err
	}

	return users, err
}

func (user *User) Delete() error {

	var err error

	err = nil
	if err = db.Where("id = ?", user.Id).Delete(&user).Error; err != nil {
		return err
	}
	return err
}
