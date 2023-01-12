package domain

type Movie struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Rating   float64 `json:"rating"`
}