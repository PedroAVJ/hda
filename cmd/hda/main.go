package main

import (
	"log"

	"github.com/PedroAVJ/hda/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handlers.SetupRoutes(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
