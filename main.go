package main

import (
	"database/sql"
	"log"

	db2 "github.com/Sotnasjeff/hexagonal-architecture-api/adapters/db"
	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")
	productDbAdapter := db2.NewProductDB(db)
	productService := app.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Example", 30)
	if err != nil {
		log.Fatalf(err.Error())
	}

	productService.Enable(product)

}
