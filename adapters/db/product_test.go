package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Sotnasjeff/hexagonal-architecture-api/adapters/db"
	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)

}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()

}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("whatever","Product Test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer DB.Close()
	productDB := db.NewProductDB(DB)

	product, err := productDB.Get("whatever")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer DB.Close()
	productDB := db.NewProductDB(DB)

	product := app.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())

}
