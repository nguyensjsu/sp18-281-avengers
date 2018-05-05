package main
import (
	"encoding/json"
	"github.com/unrolled/render"
	"fmt"
	"net/http"
	"log"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/gorilla/mux"

)

type burgerData struct{
	Title string `json:"title"`
	Price string `json:"price"`
	Description string `json:"description"`
	Code string `json:"code"`
}

	//Redis connection String
	var conn,err= redis.Dial("tcp", "54.183.231.127:6379")
	//var conn,err= redis.Dial("tcp", "localhost:6379")

	//Test Redis connection
	func Connection(){
		err := conn.Cmd("AUTH", "mypass") 
		
		fmt.Println(err)
	

		fmt.Println("Redis is now connected",conn)
		defer conn.Close()
	}

	
	//list of all the crud handlers
	func initRoutes(mx *mux.Router,formatter *render.Render){

	mx.HandleFunc("/ping",ping(formatter)).Methods("GET")
	}

	//Ping handler
	func ping(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			fmt.Println("Server running on port 30035")
			formatter.JSON(w, http.StatusOK, struct{ Burger string }{"Welcome to Counter Burger!"})
		}
	}
	func main(){
		Connection()
		formatter := render.New(render.Options{
			IndentJSON: true,
		})	
		mx := mux.NewRouter()
		initRoutes(mx,formatter)
		log.Fatal(http.ListenAndServe(":3305", mx))
		
	}