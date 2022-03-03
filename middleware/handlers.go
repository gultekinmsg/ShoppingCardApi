package middleware

import (
	"encoding/json"
	"net/http"
	"shoppingCard/db"
	"shoppingCard/utils"
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.ListItems(utils.GetQuery(r, "type"))
	err := json.NewEncoder(w).Encode(items)
	utils.CheckError(err)
}
func AddItemToBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	didSuccess := db.AddToBasket(utils.GetQuery(r, "name"))
	response := utils.Response(didSuccess)
	err := json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}
func ClearBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	didSuccess := db.ClearBasket()
	response := utils.Response(didSuccess)
	err := json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}
func ChangeBasketItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	didSuccess := db.ChangeBasket(utils.GetQuery(r, "type"), utils.GetQuery(r, "name"))
	response := utils.Response(didSuccess)
	err := json.NewEncoder(w).Encode(response)
	utils.CheckError(err)
}
