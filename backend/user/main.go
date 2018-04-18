package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	//tpl.ExecuteTemplate(w, "index.html", nil)
	template.Must(template.ParseFiles("sample.html")).Execute(w, nil)
}

func landing(w http.ResponseWriter, r *http.Request) {
	//tpl.ExecuteTemplate(w, "index.html", nil)
	template.Must(template.ParseFiles("landingpage.ejs.html")).Execute(w, nil)
}

//Employee type
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

/*
var (
	currentEmployeeID int = 2
)
*/

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "54.193.30.230:6379")
	HandleError(err)
	c.Do("AUTH", "mypass")

	return c
}

// CreatePost creates a blog post.
func initial(e Employee, num int) {

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(e)

	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "Employee:"+strconv.Itoa(e.ID), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}

func CreatePost(e Employee) {

	//currentEmployeeID++

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(e)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "Employee:"+strconv.Itoa(e.ID), b)
	HandleError(err)

	fmt.Println("GET ", reply)

}

func FindAll() Employees {

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "Employee:*")
	HandleError(err)

	var employees Employees

	for _, k := range keys.([]interface{}) {
		var employee Employee

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &employee); err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	return employees
}

func FindEmployee(id int) Employee {
	var employee Employee

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "Employee:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")
	if err = json.Unmarshal(reply.([]byte), &employee); err != nil {
		panic(err)
	}
	return employee
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
func landingpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func getThisEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r) //Get params
	id, err := strconv.Atoi(params["id"])
	fmt.Println(id)
	HandleError(err)
	employee := FindEmployee(id)
	json.NewEncoder(w).Encode(employee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	/*
		http.Redirect(w, r, "/", 302)
		r.ParseForm()
		var Firstname = r.PostFormValue("firstname")
		var Lastname = r.PostFormValue("lastname")
		var Gender = r.PostFormValue("gender")
		var Age = r.PostFormValue("age")
		IntAge, _ := strconv.Atoi(Age)

		var ID = r.PostFormValue("id")
		IntID, _ := strconv.Atoi(ID)
		var Salary = r.PostFormValue("salary")
		IntSalary, _ := strconv.Atoi(Salary)
		var employee = Employee{Firstname: Firstname, Lastname: Lastname, Gender: Gender, Age: IntAge, ID: IntID, Salary: IntSalary}
	*/

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var employee Employee

	if err := json.Unmarshal(body, &employee); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	json.NewEncoder(w).Encode(employee)
	CreatePost(employee)

	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}

	//http.FileServer(http.Dir("index.html"))

}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var employee Employee

	if err := json.Unmarshal(body, &employee); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	//fmt.Printf(employee.Firstname)

	CreatePost(employee)

	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r) //Get params
	id, err := strconv.Atoi(params["id"])
	HandleError(err)
	c := RedisConnect()
	defer c.Close()

	c.Do("DEL", "Employee:"+strconv.Itoa(id))

	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func main() {

	r := mux.NewRouter()
	initial(Employee{Firstname: "John", Lastname: "Smith", Gender: "male", Age: 27, ID: 1, Salary: 60000}, 1)
	initial(Employee{Firstname: "Steve", Lastname: "Britton", Gender: "male", Age: 30, ID: 2, Salary: 70000}, 2)

	r.HandleFunc("/", index)
	r.HandleFunc("/landing", landing)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))

	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", getThisEmployee).Methods("GET")
	r.HandleFunc("/employee", createEmployee).Methods("POST")

	r.HandleFunc("/employee/update/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employee/delete/{id}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
