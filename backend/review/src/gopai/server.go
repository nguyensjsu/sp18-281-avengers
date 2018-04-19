package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

var nodeELB = "http://iproject-elb-riak-1145643392.us-east-1.elb.amazonaws.com:80"

var debug = true

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

// Create a new client
func NewClient(server string) *Client {
	return &Client{
		Endpoint: server,
		Client:   &http.Client{Transport: tr},
	}
}

// Create a new server
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

// Ping the API to check if its working.
func (c *Client) Ping() (string, error) {
	resp, err := c.Get(c.Endpoint + "/ping")

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
// Create a new Review
func (c *Client) CreateReview(key, reqbody string) (Review, error) {
	var rev_nil = Review{}

	resp, err := c.Post(c.Endpoint+"/buckets/Review/keys/"+key+"?returnbody=true",
		"application/json", strings.NewReader(reqbody))

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return rev_nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] POST: " + c.Endpoint + "/buckets/Review/keys/" + key + "?returnbody=true => " + string(body))
	}

	var rev Review

	err = json.Unmarshal(body, &rev)

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return rev_nil, err
	}
	return rev, err
}

// Initialize our server and test ping.
func init() {
	c := NewClient(nodeELB)
	msg, err := c.Ping()

	if err != nil {
		fmt.Println("[INIT DEBUG] " + err.Error())
	} else {
		fmt.Println("Riak Ping Server: ", msg)
	}
}

// Initializing routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/review", newReviewHandler(formatter)).Methods("POST")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("[FAIL ON ERROR DEBUG] %s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

// Handles the ping call
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := NewClient(nodeELB)

		message, err := c.Ping()

		if message == "OK" {
			message = "Comment API is working."
		}

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		} else {
			formatter.JSON(w, http.StatusOK, message)
		}
	}
}

// Handle new review request
func newReviewHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var newReview Review
		uuid, _ := uuid.NewV4()

		decoder := json.NewDecoder(req.Body)
		// fmt.Println(decoder)

		err := decoder.Decode(&newReview)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		}

		newReview.Id = uuid.String()
		// ReviewItems := newReview.Items

		reqbody, _ := json.Marshal(newReview)

		c := NewClient(nodeELB)
		val_resp, err := c.CreateReview(uuid.String(), string(reqbody))

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, val_resp)
		}
	}
}