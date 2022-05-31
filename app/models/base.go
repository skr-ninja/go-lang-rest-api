package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var e error

var urlDSN = "user=postgres dbname=myDb password=1234 host=localhost sslmode=disable"

func init() {

	err = godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	//username := os.Getenv("db_user")
	//password := os.Getenv("db_pass")
	//dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	//dbPort := os.Getenv("db_port")
	//dbType := os.Getenv("db_type")
	//charset := os.Getenv("charset")
	//parseTime := os.Getenv("parse_time")

	//For MySql
	//dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&&parseTime=%s", username, password, dbHost, dbPort, dbName, charset, parseTime)
	//conn, err := gorm.Open(dbType, dbUri)

	// For Posgrace
	conn, err := gorm.Open(postgres.Open(urlDSN), &gorm.Config{})

	db = conn
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database on " + dbHost)

	db.AutoMigrate(&User{})

}

func GetDB() *gorm.DB {
	return db
}
