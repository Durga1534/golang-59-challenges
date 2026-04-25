package main

import (
	"log"
	"net/http"

	"github.com/Durga1534/go-bookstore/pkg/models"
	"github.com/Durga1534/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	models.Init()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5050", r))
}
