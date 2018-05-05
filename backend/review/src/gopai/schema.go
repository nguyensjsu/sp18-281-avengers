package main

import (
	"net/http"
)

type Client struct {
	Endpoint string
	*http.Client
}

type Review struct {
	Id        string  `json:"Id"`
	UserId    string  `json:"userId"`
	ProductId string  `json:"productId"`
	UserName  string  `json:"userName"`
	Comment   Comment `json:"comment"`
	// Status string `json:"status"`
}

type Comment struct {
	Date string
	Blob string
}

// Status string `json:"status"`

type Keys struct {
	Keys []string
}
