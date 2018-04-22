package "main"

import(
	"net/http"
)

type Client struct {
	Endpoint 	string
	*http.Client
}

type Cart struct {
	Id 		string `json:"id"`
	UserID 		string `json:"userId"`
	Date		string `json:"date"`

	Items  []struct {
		Name   	string  `json:"name"`
		Count  	int     `json:"count"`
		Rate   	float64 `json:"rate"`
		Amount 	float64 `json:"amount"`
	} `json:"items"`

	Total		float64 `json:"total"`
	Status		string	`json:"status"`
}

type Keys struct {
	Keys 		[]string 
}

type Order struct {
	OrderId		string	`json:"orderId"`
}
