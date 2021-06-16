package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee", Create).Methods("POST")
	r.HandleFunc("/employee/{eid}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
