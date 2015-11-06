package dbconnector

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db gorm.DB

//var db, _ = gorm.Open("mysql", "caprica:caprica@tcp(localhost:3306)/caprica?charset=utf8&parseTime=True&loc=Local")

func init() {

	mysql_host := os.Getenv("MYSQL_HOST")
	fmt.Println("mySQL host is ", mysql_host)
	db, _ = gorm.Open("mysql", "caprica:caprica@tcp("+mysql_host+")/caprica?charset=utf8&parseTime=True&loc=Local")

}
