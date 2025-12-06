package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = sdb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db = sdb

	println("Successfully connected to the database!")

	err = createProduct(&Product{Name: "Go Product", Price: 222})
	if err != nil {
		log.Fatal(err)
	}

	print("Successfully created product!")

}

func createProduct(product *Product) error {
	// Insert

	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name, product.Price)

	return err
}
