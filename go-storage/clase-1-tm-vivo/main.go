package main

import (
	"bcgo/internal/products"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Open database connection.
	databaseConfig := &mysql.Config{
		User:      "root",
		Passwd:    "gormit897",
		Addr:      "localhost:3306",
		DBName:    "storage",
		ParseTime: true,
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	// Ping database connection.
	if err = database.Ping(); err != nil {
		panic(err)
	}

	println("Database connection established")

	// Create products repository.
	var repository products.Repository = &products.MySQLRepository{
		Database: database,
	}

	// Get product.
	product, err := repository.Get(1)
	if err != nil {
		panic(err)
	}
}
