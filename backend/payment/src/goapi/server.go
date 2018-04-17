package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	_"os"
	"strings"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"github.com/mediocregopher/radix.v2/redis"
)

var debug = true

var serverElb = "http://riak-elb-1775435563.us-east-1.elb.amazonaws.com:80"


type Client struct {
	Endpoint string
	*http.Client
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}


func NewClient(server string) *Client {
	return &Client{
		Endpoint:  	server,
		Client: 	&http.Client{Transport: tr},
	}
}


func (c *Client) Ping() (string, error) {

	resp, err := c.Get(c.Endpoint + "/ping" )
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping Error!", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if debug { 
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/ping => " + string(body)) 
	}
	return string(body), nil
}

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func (c *Client) GetPing()(string, error) {
	resp, err := c.Get(c.Endpoint + "/ping")
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping failed", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	/*var response = map[string]string { }
	if err := json.Unmarshal({"status" : "ok"}, &response); 
	err != nil {
		fmt.Println("RIAK DEBUG] JSON unmarshaling failed: %s", err)
		return response, err
	}*/

	return string(body), nil
}

// Init Database Connections

func init() {

	// Get Environment Config

	
	// Riak KV Setup	
	c := NewClient(serverElb)
	msg, err := c.Ping( )
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Riak Ping Server: ", msg)		
	}

}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
}

func (c *Client) CreatePayment(key, reqbody string) (payment, error) {
   	var payment_nil = payment {}
	
	resp, err := c.Post(c.Endpoint + "/buckets/payment/keys/"+key+"?returnbody=true", 
						"application/json", strings.NewReader(reqbody) )
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return payment_nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if debug { fmt.Println("[RIAK DEBUG] POST: " + c.Endpoint + "/buckets/payment/keys/"+key+"?returnbody=true => " + string(body)) }
	var pay payment
	err1 := json.Unmarshal(body, &pay)
	_ = err1
 	fmt.Println(pay)
	return pay, err
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

func makePaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var pay payment
		uuid, _ := uuid.NewV4()
		decoder := json.NewDecoder(req.Body)

		err1 := decoder.Decode(&pay)

		if err1 != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println(err1)
			return
		}
		pay.Id = uuid.String()

		req_body, _ := json.Marshal(pay)
		c := NewClient(serverElb)
		pay_resp, err := c.CreatePayment(uuid.String(), string(req_body))

		if err != nil {
			log.Fatal(err)
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, pay_resp)
		}
	}
}

}