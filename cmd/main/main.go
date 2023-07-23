package main

import (
	"bookstore/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server at port 9091\n")
	log.Fatal(http.ListenAndServe("localhost:9091", r))
}
