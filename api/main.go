package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fesiqp/jwtauth/api/handlers"
	"github.com/fesiqp/jwtauth/api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbDialect string = "mysql"
	dbString  string = "jwtauth:jwtauth@(db:3306)/jwtauthdb?charset=utf8&parseTime=True&loc=Local"
)

var port string

func init() {
	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":4000"
	}
}

func main() {
	logger := log.New(os.Stdout, "[APP]   ", log.LstdFlags|log.Lshortfile)

	db, err := models.NewDB(dbDialect, dbString)
	if err != nil {
		logger.Fatal(err)
	}
	db.InitSchema()

	handlrs := handlers.New(db, logger)

	router := NewRouter(handlrs)

	err = http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal(err)
	}
}
