package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load("./2/.env"); err != nil {
		panic(err)
	}

	db := connectionSQL()

	fmt.Println("dns db connection:", db)
}

func connectionSQL() string {
	// vars
	user := os.Getenv("DB_USER")
	pswd := os.Getenv("DB_PSWD")
	host := os.Getenv("DB_HOST") //127.0.0.1
	port := os.Getenv("DB_PORT") //3306
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pswd, host, port, name)
}