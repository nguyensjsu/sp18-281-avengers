package main

import (
	"fmt"
	"math"
	"time"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
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

// Ping the API to check if its working.
func (c *Client) Ping() (string, error) {
	resp, err := c.Get(c.Endpoint + "/ping")

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping Error!", err
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// Create a new order
func (c *Client) CreateOrder(key, reqbody string) (Cart, error) {
	var ord_nil = Cart{}

	resp, err := c.Post(c.Endpoint + "/buckets/Orders/keys/" + key + "?returnbody=true",
										"application/json", strings.NewReader(reqbody))

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var place Cart

	err = json.Unmarshal(body, &place)

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}
	return place, err
}

// Update order for updating completing order.
func (c *Client) UpdateOrder(cartEdit Cart) (Cart, error) {
	var ord_nil = Cart {}
	reqbody, _ := json.Marshal(cartEdit)

	req_body := string(reqbody)

	req, _  := http.NewRequest("PUT", c.Endpoint + "/buckets/Orders/keys/"+ cartEdit.Id +"?returnbody=true", strings.NewReader(req_body) )
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	var ord Cart

	err = json.Unmarshal(body, &ord)
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}
	return ord, err
}

// Get keys of all objects stored in database.
func (c *Client) GetKeys() ([]string, error) {
	var keys_nil [] string

	resp, err := c.Get(c.Endpoint + "/buckets/Orders/keys?keys=true")
	
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return keys_nil, err
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	var all_keys Keys
	err = json.Unmarshal(body, &all_keys)
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return keys_nil, err
	}
 
	return all_keys.Keys, err
}



// View order of specific key
func (c *Client) GetOrder(key string) (Cart) {

	var ord_nil = Cart {}
	resp, err := c.Get(c.Endpoint + "/buckets/Orders/keys/" + key)
		
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var cart Cart
	err = json.Unmarshal(body, &cart)
	if err != nil {
		fmt.Println("RIAK DEBUG] JSON unmarshaling failed: %s", err)
		return ord_nil
	}
	fmt.Println("[TEST] ", cart)	


	var ord_nil = Cart {}
	resp, err := c.Get(c.Endpoint + "/buckets/Orders/keys/" + key )

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	var ord = Cart {}
	
	if err := json.Unmarshal(body, &ord); err != nil {
		fmt.Println("RIAK DEBUG] JSON unmarshaling failed: %s", err)
		return ord_nil
	}
	return ord
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
	mx.HandleFunc("/order", createOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/order/{id}", updateCartHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/view/{id}", getOrderHandler(formatter)).Methods("GET")
	mx.HandleFunc("/history/{id}", viewCartHandler(formatter)).Methods("GET")
	mx.HandleFunc("/clearCart/{id}", clearCartHandler(formatter)).Methods("DELETE")
	
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

unc pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}
		c := NewClient(nodeELB)

		message, err := c.Ping()

		if message == "OK" {
			message = "Cart API is working."
		}

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		} else {
			formatter.JSON(w, http.StatusOK, message)
		}
	}
}

func createOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}
		var newCart Cart
		uuid, _ := uuid.NewV4()
		
		decoder := json.NewDecoder(req.Body)

		err := decoder.Decode(&newCart)
		fmt.Println("*************")
		fmt.Println(newCart)
		fmt.Println("*************")
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		}

		newCart.Id = uuid.String()
		newCart.Date = time.Now().Local().Format("2006/01/02")
		newCart.Status = "IN CART"
		cartItems := newCart.Items

		var totalAmount float64

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].Amount = calculateAmount(cartItems[i].Count, cartItems[i].Rate)
			totalAmount += cartItems[i].Amount
		}

		totalAmount = math.Ceil(totalAmount * 100) / 100
		newCart.Total = totalAmount

		reqbody, _ := json.Marshal(newCart)

		c := NewClient(nodeELB)
		val_resp, err := c.CreateOrder(uuid.String(), string(reqbody))

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, val_resp)
		}
	}
}

func updateCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}

		params := mux.Vars(req)
		var uuid string = params["id"]

		if uuid == "" {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request. User ID Missing.")
		} else {
			var newCart Cart
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&newCart)

			if err != nil {
				ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
				fmt.Println("[HANDLER DEBUG] ", err.Error())
				return
			}

			var totalAmount float64

			cartItems := newCart.Items

			for i := 0; i < len(cartItems); i++ {
				cartItems[i].Amount = calculateAmount(cartItems[i].Count, cartItems[i].Rate)
				totalAmount += cartItems[i].Amount
			}
			totalAmount = math.Ceil(totalAmount * 100) / 100

			newCart.Total = totalAmount
			newCart.Id = uuid
			newCart.Date = time.Now().Local().Format("2006/01/02")
			
			c := NewClient(nodeELB)
			reqbody, _ := json.Marshal(newCart)
			val_resp, err := c.CreateOrder(uuid, string(reqbody))

			if err != nil {
				fmt.Println("[HANDLER DEBUG] ", err.Error())
				formatter.JSON(w, http.StatusBadRequest, err)
			} else {
				formatter.JSON(w, http.StatusOK, val_resp)
			}
		}
	}
}

// To view our order
func getOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
		params := mux.Vars(req)
		var uuid string = params["id"]
		// fmt.Println( "Order Params ID: ", uuid )

		if uuid == ""  {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request. Order ID Missing.")
		} else {

			c := NewClient(nodeELB)
			keys, _ := c.GetKeys();
			var prev_cart Cart
			for _ , item := range keys {
				cart := c.GetOrder(item)
				if cart.UserID == uuid  && cart.Status == "IN CART"{
					prev_cart = cart
				}
			}
			if prev_cart.Id == "" {
				formatter.JSON(w, http.StatusNoContent, nil)
			} else {
				formatter.JSON(w, http.StatusOK, prev_cart)
			}
		}
	}
}

func viewCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var uid string = params["id"]

		if uid == "" {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request. User ID Missing.")
		} else {
			c := NewClient(nodeELB)
			
			cart_keys, err := c.GetKeys()
			cart_list := []Cart{}
			for _ , item := range cart_keys {
				if(c.GetOrder(item).UserID == uid) {
					if(c.GetOrder(item).Status != "CLEARED") {
						cart_list = append(cart_list, c.GetOrder(item))
					}
				}
			}

			if err != nil {
				fmt.Println("[HANDLER DEBUG] ", err.Error())
				formatter.JSON(w, http.StatusBadRequest, err)
			} else {
				formatter.JSON(w, http.StatusOK, cart_list)
			}
		}
		
	}
}

func calculateAmount(count int, rate float64) float64{
	total := float64(count) * rate
	total = math.Ceil(total * 100) / 100
	return total
}

// Delete current order
func clearCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
/*		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}*/

		params := mux.Vars(req)
		var uuid string = params["id"]

		if uuid == "" {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request. Order ID Missing.")
		} else {
			c := NewClient(nodeELB)

			ord := c.GetOrder(uuid)

			if ord.Id == "" {
				formatter.JSON(w, http.StatusBadRequest, "")
			}
			if ord.Status == "IN CART" {
				err := c.DeleteOrder(uuid)

				if err != nil {
					fmt.Println("[HANDLER DEBUG] ", err.Error())
					formatter.JSON(w, http.StatusBadRequest, err)
				} else {
					formatter.JSON(w, http.StatusOK, "Cart cleared successfully")
				}

			} else {
				formatter.JSON(w, http.StatusOK, "Can't perform this action.")
			}
		}
	}
}
