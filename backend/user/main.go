package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

//Book Struct (Model)
var client *redis.Client
var templates *template.Template

type Employee struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	ID        int    `json:"id"`
	Salary    int    `json:"salary"`
}

// A list of employees
type Employees []Employee

//Get all books
func getEmployees(w http.ResponseWriter, r *http.Request) {

}

//Get a single book
func getThisEmployee(w http.ResponseWriter, r *http.Request) {

}

//Create a new book
func createEmployee(w http.ResponseWriter, r *http.Request) {

}

//updateBook
func updateEmployee(w http.ResponseWriter, r *http.Request) {

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))

	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", getThisEmployee).Methods("GET")
	r.HandleFunc("/employee", createEmployee).Methods("POST")

	r.HandleFunc("/employee/update/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employee/delete/{id}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
