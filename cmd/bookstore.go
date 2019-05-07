package main

import (
	"log"

	"github.com/alvadorn/bookstore/internal/config"
	books "github.com/alvadorn/bookstore/pkg/books/handler"
	"github.com/jinzhu/gorm"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var db *gorm.DB
var err error

func init() {
	log.Println("Initializing DB")

	db, err = config.ConnectDB()

	if err != nil {
		log.Println("Error on connection")
		log.Println(err)
		err = nil
	}

	config.AutoMigrate(db)
}

func main() {

	router := routing.New()

	v1 := router.Group("/v1")

	books.Handler(v1.Group("/books"))

	log.Println("Starting server on port 3000")

	panic(fasthttp.ListenAndServe(":3000", router.HandleRequest))

}
