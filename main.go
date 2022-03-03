package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"shoppingCard/router"
	"shoppingCard/utils"
)

func main() {
	err := godotenv.Load(".env")
	utils.CheckError(err)
	r := router.Router()
	err = http.ListenAndServe(os.Getenv("PORT"), r)
	utils.CheckError(err)
}
