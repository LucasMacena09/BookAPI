package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/LucasMacena09/BookAPI/handlers"
)

func SetupRoutes(app *fiber.App, repo *handlers.Repository) {
    api := app.Group("/api")
    api.Post("/create_books", repo.CreateBook)
    api.Delete("/delete_book/:id", repo.DeleteBook)
    api.Get("/get_books", repo.GetBooks)
    api.Get("/get_book/:id", repo.GetBookByID)
}