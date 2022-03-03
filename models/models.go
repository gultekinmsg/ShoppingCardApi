package models

type Item struct {
	NAME           string `json:"name"`
	PRICE          int    `json:"price"`
	INBASKET       bool   `json:"inbasket"`
	BASKETQUANTITY int    `json:"basketquantity"`
	BASKETPRICE    int    `json:"basketprice"`
}
