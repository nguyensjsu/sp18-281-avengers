package main

type payment struct {
	Id             	string 	
	UserId			string
	CardNumber   	int    	
	OrderId			string
	CardType 		string	    
	CardHolderName 	string	
	Amount			int
}

type keys struct {
	Keys 			[]string 
}