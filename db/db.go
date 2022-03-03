package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"shoppingCard/models"
	"shoppingCard/utils"
)

func createConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	utils.CheckError(err)
	err = db.Ping()
	utils.CheckError(err)
	return db
}
func ListItems(whichList string) []models.Item {
	db := createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	selectStatement := `select "Item", "Price","InBasket","BasketQuantity","BasketPrice" from shopping."ShoppingItems" order by "Price"`
	result, err := db.Query(selectStatement)
	utils.CheckError(err)
	return utils.SqlToObject(result, whichList)
}
func AddToBasket(name string) bool {
	db := createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	updateStatement := `update shopping."ShoppingItems" set "InBasket" = true where "Item" = $1`
	_, err := db.Exec(updateStatement, name)
	utils.CheckError(err)

	updateStatement = `update shopping."ShoppingItems" set "BasketPrice" = "Price" where "Item" = $1`
	_, err = db.Exec(updateStatement, name)
	utils.CheckError(err)

	updateStatement = `update shopping."ShoppingItems" set "BasketQuantity" = 1 where "Item" = $1`
	_, err = db.Exec(updateStatement, name)
	utils.CheckError(err)

	if err == nil {
		return true
	} else {
		return false
	}
}
func ClearBasket() bool {
	db := createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	updateStatement := `update shopping."ShoppingItems" set "InBasket" = false`
	_, err := db.Exec(updateStatement)
	utils.CheckError(err)

	updateStatement = `update shopping."ShoppingItems" set "BasketPrice" = 0`
	_, err = db.Exec(updateStatement)
	utils.CheckError(err)

	updateStatement = `update shopping."ShoppingItems" set "BasketQuantity" = 0`
	_, err = db.Exec(updateStatement)
	utils.CheckError(err)
	if err == nil {
		return true
	} else {
		return false
	}
}
func ChangeBasket(operation, name string) bool {
	db := createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if operation == utils.REMOVE {
		updateStatement := `update shopping."ShoppingItems" set "InBasket" = false where "Item" = $1`
		_, err := db.Exec(updateStatement, name)
		utils.CheckError(err)

		updateStatement = `update shopping."ShoppingItems" set "BasketPrice" = 0 where "Item" = $1`
		_, err = db.Exec(updateStatement, name)
		utils.CheckError(err)

		updateStatement = `update shopping."ShoppingItems" set "BasketQuantity" = 0 where "Item" = $1`
		_, err = db.Exec(updateStatement, name)
		utils.CheckError(err)

		if err == nil {
			return true
		} else {
			return false
		}
	} else if operation == utils.INCREASE {
		selectStatement1 := `select "Item", "Price","InBasket","BasketQuantity","BasketPrice" from shopping."ShoppingItems"`
		result1, err := db.Query(selectStatement1)
		utils.CheckError(err)
		items1 := utils.SqlToObject(result1, utils.BASKET)

		basketQuantity := items1[0].BASKETQUANTITY + 1
		updateStatement := `update shopping."ShoppingItems" set "BasketQuantity" = $1 where "Item" = $2`
		_, err = db.Exec(updateStatement, basketQuantity, name)
		utils.CheckError(err)

		basketPrice := items1[0].PRICE * basketQuantity
		updateStatement = `update shopping."ShoppingItems" set "BasketPrice" = $1 where "Item" = $2`
		_, err = db.Exec(updateStatement, basketPrice, name)
		utils.CheckError(err)

		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		selectStatement1 := `select "Item", "Price","InBasket","BasketQuantity","BasketPrice" from shopping."ShoppingItems"`
		result1, err := db.Query(selectStatement1)
		utils.CheckError(err)
		items1 := utils.SqlToObject(result1, utils.BASKET)

		basketQuantity := items1[0].BASKETQUANTITY - 1
		updateStatement := `update shopping."ShoppingItems" set "BasketQuantity" = $1 where "Item" = $2`
		_, err = db.Exec(updateStatement, basketQuantity, name)
		utils.CheckError(err)

		basketPrice := items1[0].PRICE * basketQuantity
		updateStatement = `update shopping."ShoppingItems" set "BasketPrice" = $1 where "Item" = $2`
		_, err = db.Exec(updateStatement, basketPrice, name)
		utils.CheckError(err)

		if err == nil {
			return true
		} else {
			return false
		}
	}
}
