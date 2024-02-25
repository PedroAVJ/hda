package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/PedroAVJ/hda/internal/handlers"
)

func main() {
	r := gin.Default()

	handlers.SetupRoutes(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
