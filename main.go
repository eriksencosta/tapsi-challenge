package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

var port string = ":7100"
var method string = "POST"
var path string = "/move"


func main() {
	router := mux.NewRouter()
	router.HandleFunc( path , getMovement).Methods(method)
	log.Fatal(http.ListenAndServe(port, router))
}
