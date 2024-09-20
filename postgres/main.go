package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"my-module/models"

	_ "github.com/lib/pq"
)



func CreateTable(db *sql.DB) {

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		product VARCHAR(100) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	db.Query(createTableSQL)

	log.Println("Products table created successfully")

}

func InsertTable(db *sql.DB, p models.Product) {

	// SQL statement for inserting a product
	sqlStatement := `
    INSERT INTO products (product, price)
    VALUES ($1, $2)
    RETURNING id`

	// Execute the SQL statement
	var id int
	err := db.QueryRow(sqlStatement, p.Product, p.Price).Scan(&id)
	if err != nil {
		fmt.Printf("failed to insert product: %v", err)
	}

	fmt.Printf("Inserted product with ID %d\n", id)
}

func GetRows(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal("error: ", err)
		return nil, err
	}
	return rows, nil

}
