package main

import (
	"database/sql"
	"fmt"
)

type Subcategory struct {
	SubcategoryID   int    `json:"SubcategoryID"`
	SubcategoryName string `json:"SubcategoryName"`
	Products        Product
}

func (subcategory *Subcategory) getVarient(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT subcategory.subcategoryID,subcategory.subcategoryName, products.productID,products.productName,products.productPrice,products.subcategoryID from subcategory,products where subcategoryName='%s' and subcategory.subcategoryID=products.subcategoryID", subcategory.SubcategoryName)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&subcategory.SubcategoryID, &subcategory.SubcategoryName, &subcategory.Products.ProductID,
		&subcategory.Products.ProductName, &subcategory.Products.ProductPrice)
}

func (subcategory *Subcategory) updateVarient(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE subcategory SET subcategoryName='%s'WHERE subcategoryID=%d", subcategory.SubcategoryName, subcategory.SubcategoryID)
	_, err := db.Exec(statement)
	return err
}

func (subcategory *Subcategory) deleteVarient(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM subcategory WHERE subcategoryID=%d", subcategory.SubcategoryID)
	_, err := db.Exec(statement)
	return err
}

func (subcategory *Subcategory) createVarient(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO subcategory (subcategoryID,subcategoryName) VALUES('%d','%s')",
		subcategory.SubcategoryID, subcategory.SubcategoryName)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_VarID()").Scan(&varient.VarID)

	if err != nil {
		return err
	}

	return nil
}
func getVarients(db *sql.DB, start, count int) ([]Subcategory, error) {
	statement := fmt.Sprintf("SELECT *  FROM subcategory LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Subcategorys := []Subcategory{}

	for rows.Next() {
		var subcategory Subcategory
		if err := rows.Scan(&subcategory.SubcategoryID, &subcategory.SubcategoryName); err != nil {
			return nil, err
		}
		Subcategorys = append(Subcategorys, subcategory)
	}

	return Subcategorys, nil
}
