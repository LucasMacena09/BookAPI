package handlers

import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/LucasMacena09/BookAPI/models"
    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
    var book models.Book
    
    if err := c.BodyParser(&book); err != nil {
        return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid request"})
    }

    if book.Author == nil || book.Title == nil || book.Publisher == nil ||
       *book.Author == "" || *book.Title == "" || *book.Publisher == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "All fields (author, title, publisher) are required"})
    }

    if err := r.DB.Create(&book).Error; err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not create book"})
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Book added"})
}


func (r *Repository) DeleteBook(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID cannot be empty"})
    }
    if err := r.DB.Delete(&models.Book{}, id).Error; err != nil {
        return c.Status(http.StatusBadGateway).JSON(fiber.Map{"message": "Could not delete book"})
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
    var books []models.Book
    if err := r.DB.Find(&books).Error; err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not fetch books"})
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Books fetched successfully", "data": books})
}

func (r *Repository) GetBookByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if id == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID cannot be empty"})
    }
    if err := r.DB.First(&book, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Book not found"})
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book fetched successfully", "data": book})
}