package main

import (
	"log"
	"os"

	"github.com/PedroAVJ/hda/pkg/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ToDo struct {
	ID        uint `gorm:"primaryKey"`
	Task      string
	Completed bool
}

func main() {
	// Load environment variables from file.
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Connect to PlanetScale database using DSN environment variable.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	if err := db.AutoMigrate(&ToDo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	app := echo.New()
	app.GET("/", handlers.HomeHandler)
	app.Logger.Fatal(app.Start(":4000"))
}
