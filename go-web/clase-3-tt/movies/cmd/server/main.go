package main

import (
	"encoding/json"
	"log"
	"os"
	"rest/cmd/server/handlers"
	"rest/internal/domain"
	"rest/internal/movies"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// instances
	// db, err := connectDB("./movies.json")
	db, err := connectDB(os.Getenv("DB_FILE"))
	if err != nil {
		panic(err)
	}
	// internal
	rp := movies.NewRepository(&db)
	sv := movies.NewService(rp)

	
	// server [exposure]
	server := gin.Default()

	movies := server.Group("/movies")
	h := handlers.NewMovies(sv)
	movies.GET("", h.Get())
	movies.GET("/:id", h.GetByID())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func connectDB(filename string) ([]domain.Movie, error) {
	var movies []domain.Movie

	// reader
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decoder
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&movies); err != nil {
		return nil, err
	}
	
	return movies, nil
}