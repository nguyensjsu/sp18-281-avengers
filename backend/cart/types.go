package "main"

type Client struct {
	Endpoint 	string
	*http.Client
}


