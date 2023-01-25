package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// env
	os.Setenv("GOOGLE_AUTH_FILE", "./1/credentials.json")

	// app
	auth, err := IsAuthenticated()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("auth:", auth)
}

type credentials struct {
	Token	string `json:"token"`
}

func IsAuthenticated() (bool, error) {
	filename := os.Getenv("GOOGLE_AUTH_FILE")
	fmt.Println("key:", filename)

	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var cred credentials
	if err := decoder.Decode(&cred); err != nil {
		return false, err
	}

	// process
	// return cred.Token == "123", nil
	if cred.Token == "123" {
		return true, nil
	}
	return false, nil
}