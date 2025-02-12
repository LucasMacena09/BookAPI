package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api = app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("get_books/:id", r.GetBook)
	api.Put("put_books/:id", r.)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r := Repository {
		DB: db, 
	}

	app := fiber.New()
	r.SetupRoute(app)
	app.Listen(":8080")
}