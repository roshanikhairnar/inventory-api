package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ProductID     int    `json:"ProductID"`
	ProductName   string `json:"ProductName"`
	ProductPrice  int    `json:"ProductPrice"`
	SubcategoryID int    `json:"SubcategoryID"`
}

func getproductswithvar(db *sql.DB, start, count int) ([]Product, error) {
	statement := fmt.Sprintf("select product.ProductName, product.ProductID products LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Varientswithproduct := []Product{}

	for rows.Next() {
		//var varient Varient
		var product Product
		if err := rows.Scan(&product.ProductName, &product.ProductID); err != nil {
			return nil, err
		}
		Varientswithproduct = append(Varientswithproduct, product)
	}

	return Varientswithproduct, nil
}

func (product *Product) getProduct(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM products where productName='%s'", product.ProductName)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&product.ProductID, &product.ProductName, &product.SubcategoryID)
}

func (product *Product) updateProduct(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE products SET ProductName='%s'WHERE ProductID=%d", product.ProductName, product.ProductID)
	_, err := db.Exec(statement)
	return err
}

func (product *Product) deleteProduct(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM products WHERE productID=%d", product.ProductID)
	_, err := db.Exec(statement)
	return err
}

func (product *Product) createProduct(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO products VALUES('%d','%s','%d')", product.ProductID, product.ProductName, product.SubcategoryID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_ProductID()").Scan(&Product.ProductID)

	if err != nil {
		return err
	}

	return nil
}
func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	statement := fmt.Sprintf("SELECT productID, productName, subcategoryID FROM products LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Products := []Product{}

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.SubcategoryID); err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}

	return Products, nil
}
