package db

import (
	"database/sql"

	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(id string) (app.ProductInterface, error) {
	var product app.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int
	p.db.QueryRow("SELECT idFROM products WHERE id = ?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDB) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products(id, name, price, status) VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product app.ProductInterface) (app.ProductInterface, error) {
	_, err := p.db.Exec(`UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?`,
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetId())

	if err != nil {
		return nil, err
	}

	return product, nil
}
