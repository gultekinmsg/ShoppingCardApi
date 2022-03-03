package utils

import (
	"database/sql"
	"log"
	"net/http"
	"shoppingCard/models"
)

const (
	REMOVE   = "REMOVE"
	DECREASE = "DECREASE"
	INCREASE = "INCREASE"
	ALL      = "ALL"
	BASKET   = "BASKET"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SqlToObject(result *sql.Rows, whichList string) []models.Item {
	var list []models.Item
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result)
	for result.Next() {
		var item models.Item
		err := result.Scan(&item.NAME, &item.PRICE, &item.INBASKET, &item.BASKETQUANTITY, &item.BASKETPRICE)
		CheckError(err)
		if whichList == BASKET {
			if item.INBASKET {
				list = append(list, item)
			}
		} else {
			list = append(list, item)
		}
	}
	return list
}
func GetQuery(r *http.Request, what string) string {
	keys, ok := r.URL.Query()[what]
	if !ok {
		log.Fatalf("Cannot retrieve " + what + " from request")
	}
	return keys[0]
}
func Response(didSuccessful bool) string {
	if didSuccessful {
		return "200"
	} else {
		return "500"
	}
}
