package main

import (
	"log"
	"rest/cmd/routes"
	"rest/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	// instances
	db := []domain.Website{
		{ID: 1, URL: "https://www.amazon.com", Host: "amazon", Category: "e-commerce", Protected: false},
		{ID: 2, URL: "https://www.primevideo.com", Host: "amazon", Category: "streaming", Protected: true},
		{ID: 3, URL: "https://www.netflix.com", Host: "netflix", Category: "streaming", Protected: false},
	}
	// rp := website.NewRepository(&db, 3)
	// sv := website.NewService(rp)
	
	// app 
	// websites, _ := sv.Get()
	// fmt.Println("- websites:", websites)
	
	// ws, err := sv.Create("https://www.music.com", "amazon", "music", false)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println("- website created:", ws)

	// exposure [std io]
	en := gin.Default()
	rt := routes.NewRouter(en, &db)
	rt.SetRoutes()

	if err := en.Run(); err != nil {
		log.Fatal(err)
	}
}