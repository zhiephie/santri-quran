package main

import (
	"fmt"
	"log"

	"github.com/santri-quran/database"
	"github.com/santri-quran/quran"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api") // /api

	v1 := api.Group("/v1")

	v1.Get("/surah", quran.GetSurahs)
	v1.Get("/surah/:id", quran.GetSurah)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "quran.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// database.DBConn.AutoMigrate(&quran.Quran{})
	// database.DBConn.AutoMigrate(&quran.Surah{})
	// fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
