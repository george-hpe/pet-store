package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// repository represent the repository model
type repository struct {
	db *sql.DB
}

// Product represents a product in the store.
type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	URL      string `json:"url"`
	Status   string `json:"status"`
}

// Repository represent the repositories.
type Repository interface {
	AddProduct(p *Product) (string, error)
	ListProduct() ([]Product, error)
	Close() error
}

// NewRepository will create a variable that represent the Repository object.
func NewRepository() Repository {

	fileName := "database/sqlite-database.db"

	// Create SQLite file.
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create SQLite file: %v", err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Create the tables.
	query, err := ioutil.ReadFile("database/db.sql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}
	if _, err := db.Exec(string(query)); err != nil {
		log.Fatalf("Failed to set up tables: %v", err)
	}
	return &repository{db: db}
}

// AddProduct add a new product to the store.
func (r *repository) AddProduct(p *Product) (string, error) {

	stmt, err := r.db.Prepare("INSERT INTO Products(prod_id, prod_name, prod_category, prod_photo_url, status) values(?,?,?,?,?)")
	if err != nil {
		return "", fmt.Errorf("query could not be prepared")
	}

	_, err = stmt.Exec(p.ID, p.Name, p.Category, p.URL, p.Status)
	if err != nil {
		return "", fmt.Errorf("could not execute the query: %v", err)
	}

	return p.ID, nil
}

// ListProduct fetch all products from the store.
func (r *repository) ListProduct() ([]Product, error) {

	rows, err := r.db.Query("SELECT * FROM Products")
	if err != nil {
		return nil, fmt.Errorf("failed to execute the select query")
	}

	var res []Product

	for rows.Next() {
		var p Product
		if err = rows.Scan(&p.ID, &p.Name, &p.Category, &p.URL, &p.Status); err == nil {
			res = append(res, p)
		}
	}
	rows.Close()
	return res, nil
}

// Close ends the db connection.
func (r *repository) Close() error {
	return r.db.Close()
}
