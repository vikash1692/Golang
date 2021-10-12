package main

import (
	"log"
	"net/http"
	"github.com/vikash1692/Golang/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}