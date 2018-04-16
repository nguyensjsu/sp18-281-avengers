package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	//"log"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	employee1 := &Employee{Firstname: "John", Lastname: "Smith", Gender: "Male", Age: "27", Salary: 60000}
	data, err := json.Marshal(&employee1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	err = client.Set("key", data, 0).Err()
	if err != nil {
		panic(err)
	}

}
