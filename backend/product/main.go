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
<<<<<<< HEAD
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
=======
	//display Burger item
	mx.HandleFunc("/displayitem",getItem(formatter)).Methods("GET") 
	//Create a new Burger item
	mx.HandleFunc("/createitem",createItem(formatter)).Methods("POST")
	mx.HandleFunc("/createitem",createItem(formatter)).Methods("OPTIONS")
	//update Burger item
	mx.HandleFunc("/item",updateItem(formatter)).Methods("PUT")
	mx.HandleFunc("/item",createItem(formatter)).Methods("OPTIONS")
	//delete Burger item
	mx.HandleFunc("/item",deleteItem(formatter)).Methods("DELETE")
	mx.HandleFunc("/item",createItem(formatter)).Methods("OPTIONS")
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


		//Display the burger
	func getItem(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request){

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			if req.Method == "OPTIONS"{
				return
			}
			
			var conn,err= redis.Dial("tcp", "54.183.231.127:6379")
			err1 := conn.Cmd("AUTH", "mypass") 
		
			fmt.Println(err1)
			//conn,err:= redis.Dial("tcp", "localhost:6379")
			if err != nil {
                log.Fatal(err)
			}
			defer conn.Close()
			fmt.Println("Server running on port 30035")
			allburgers:= conn.Cmd("HGETAll", "Burger")
			l,_:=allburgers.List()
			fmt.Println("List is",l)  
    		if err != nil {
        	log.Fatal(err)
    		}			
			formatter.JSON(w, http.StatusOK,l)
		}
	}

	//Create burger
	func createItem(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request){
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			if req.Method == "OPTIONS"{
				return
			}

	var m burgerData
		err=json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error in Decoding the data")
				panic(err)
			}

		fmt.Println("title is :",m.Title)
		fmt.Println("price is :",m.Price)
		fmt.Println("description is :",m.Description)

			conn,err:= redis.Dial("tcp", "54.183.231.127:6379")
			//var conn,err= redis.Dial("tcp", "localhost:6379")
			err1 := conn.Cmd("AUTH", "mypass") 
		
		fmt.Println(err1)
			if err != nil {
                log.Fatal(err)
			}
			defer conn.Close()
			fmt.Println("Value in code:",m.Code)
			result:= conn.Cmd("HMSET", "Burger",m.Code+"Code",m.Code, m.Code+"Title",m.Title,m.Code+"Price",m.Price,m.Code+"Description",m.Description)
			fmt.Println("Created New Item",result)
			var response string="Created as new Burger Recipe!"
			formatter.JSON(w, http.StatusOK, struct{ Item string `json:"Item,omitempty"` }{response})
		}
	}


	//update burger
	func updateItem(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request){
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")

			if req.Method == "OPTIONS"{
				return
			}

			var m burgerData
				err=json.NewDecoder(req.Body).Decode(&m)
				if err != nil {
					fmt.Println("Error in Decoding the data")
						panic(err)
					}

			
			conn,err:= redis.Dial("tcp", "54.183.231.127:6379")
			//var conn,err= redis.Dial("tcp", "localhost:6379")
			err1 := conn.Cmd("AUTH", "mypass") 
		
		fmt.Println(err1)
			if err != nil {
                log.Fatal(err)
			}
			defer conn.Close()
			val:=conn.Cmd("HEXISTS","Burger",m.Code+"Title")		
			if val.Err != nil{
				fmt.Println("There is no so key to update",val.Err)
				panic(val.Err)
			}else {
			result:= conn.Cmd("HMSET", "Burger",m.Code+"Price",m.Price )
			fmt.Println("After Price Update",result)
			var response string="Updated Price of Burger"
			formatter.JSON(w, http.StatusOK, struct{ Item string `json:"Item,omitempty"` }{response})
			}
		}
	}


	//Delete Burger
	func deleteItem(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request){

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			if req.Method == "OPTIONS"{
				return
			}
			var m burgerData
				err=json.NewDecoder(req.Body).Decode(&m)
				if err != nil {
					fmt.Println("Error in Decoding the data")
						panic(err)
					}

			
			conn,err:= redis.Dial("tcp", "54.183.231.127:6379")
			//conn,err= redis.Dial("tcp", "localhost:6379")
			err1 := conn.Cmd("AUTH", "mypass") 
		
		fmt.Println(err1)
			if err != nil {
                log.Fatal(err)
			}
			defer conn.Close()
			val:=conn.Cmd("HEXISTS","Burger",m.Code+"Title")
			if val.Err != nil{
				fmt.Println("There is no so key to update",val.Err)
				panic(val.Err)
			}else {
			result:= conn.Cmd("HDEL", "Burger",m.Code+"Code",m.Code+"Title",m.Code+"Price",m.Code+"Description")
			fmt.Println("After item deleted",result)
			var response string="This Burger will longer be available in the Catalog"
			formatter.JSON(w, http.StatusOK, struct{ Item string `json:"Item,omitempty"` }{response})
			}
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
>>>>>>> 9c022dbd294e4f7f17fadc13d9ed18f2e304b0e5
