package main

import (
	_"fmt"
	_"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// render: for rendering JSON file
// negroni: server
// mux: request handler
// log: to log errors
// fmt: for print statement

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

func NewClient(server string) *Client {
	return &Client{
		Endpoint: server,
		Client:   &http.Client{Transport: tr},
	}
}



func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options {
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func init() {
	c := NewClient(nodeELB)
	msg, err := c.Ping()

	if err != nil {
		fmt.Println("[INIT DEBUG] " + err.Error())
	} else {
		fmt.Println("Riak Ping Server: ", msg)
	}
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", pingHandler(formatter)).Methods("GET")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("[FAIL ON ERROR DEBUG] %s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{Test string}{"Test API is live!!"})
	}
}