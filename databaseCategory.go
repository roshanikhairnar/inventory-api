package main

import (
	"database/sql"
	"fmt"
)

type Category struct {
	CategoryID  int    `json:"categoryID"`
	Name        string `json:"categoryName"`
	Subcategory Subcategory
}

func (category *Category) getCategory(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT category.categoryID,category.categoryName,subcategory.subcategoryID,subcategory.subcategoryName,products.productID,products.productName,products.productPrice,products.subcategoryID from category,subcategory,products where category.CategoryID=subcategory.categoryID and subcategory.subcategoryID=products.subcategoryID and category.categoryName='%s'", category.Name)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&category.CategoryID, &category.Name, &category.Subcategory.SubcategoryID, &category.Subcategory.SubcategoryName,
		&category.Subcategory.Products.ProductID, &category.Subcategory.Products.ProductName, &category.Subcategory.Products.ProductPrice, &category.Subcategory.Products.SubcategoryID)
}

func (category *Category) updateCategory(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE category SET categoryName='%s'WHERE categoryID=%d", category.Name, category.CategoryID)
	_, err := db.Exec(statement)
	return err
}

func (category *Category) deleteCategory(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM category WHERE categoryID=%d", category.CategoryID)
	_, err := db.Exec(statement)
	return err
}

func (category *Category) createCategory(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO category(categoryName,categoryID) VALUES('%s','%d')", category.Name, category.CategoryID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_CategoryID()").Scan(&category.CategoryID)

	if err != nil {
		return err
	}

	return nil
}

func getCategorys(db *sql.DB, start, count int) ([]Category, error) {
	statement := fmt.Sprintf("SELECT category.categoryID,category.categoryName,subcategory.subcategoryID,subcategory.subcategoryName,products.productID,products.productName,products.productPrice,products.subcategoryID from category,subcategory,products where category.CategoryID=subcategory.categoryID and subcategory.subcategoryID=products.subcategoryID LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Categorys := []Category{}

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryID, &category.Name, &category.Subcategory.SubcategoryID, &category.Subcategory.SubcategoryName,
			&category.Subcategory.Products.ProductID, &category.Subcategory.Products.ProductName, &category.Subcategory.Products.ProductPrice, &category.Subcategory.Products.SubcategoryID); err != nil {
			return nil, err
		}
		Categorys = append(Categorys, category)
	}

	return Categorys, nil
}
