package router

import (
	"github.com/gorilla/mux"
	"shoppingCard/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/list", middleware.ListItems).Methods("GET", "OPTIONS")
	router.HandleFunc("/putIn", middleware.AddItemToBasket).Methods("PUT", "OPTIONS")
	router.HandleFunc("/clear", middleware.ClearBasket).Methods("PUT", "OPTIONS")
	router.HandleFunc("/change", middleware.ChangeBasketItem).Methods("PUT", "OPTIONS")
	return router
}
