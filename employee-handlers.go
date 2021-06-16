package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee 
	Database.Find(&employees)

	json.NewEncoder(w).Encode(employees)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee 
	Database.First(&employee,mux.Vars(r)["eid"]) 

	json.NewEncoder(w).Encode(employee)
}
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee 
	Database.First(&employee,mux.Vars(r)["eid"]) 
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)

	json.NewEncoder(w).Encode(employee)
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee 
	Database.Delete	(&emp,mux.Vars(r)["eid"]) 
	

	json.NewEncoder(w).Encode("employee is deleted")
}




