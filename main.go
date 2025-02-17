package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "github.com/LucasMacena09/BookAPI/storage"
    "github.com/LucasMacena09/BookAPI/models"
    "github.com/LucasMacena09/BookAPI/handlers"
    "github.com/LucasMacena09/BookAPI/routes"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    config := &storage.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User:     os.Getenv("DB_USER"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
        DBName:   os.Getenv("DB_NAME"),
    }

    db, err := storage.NewConnection(config)
    if err != nil {
        log.Fatal("Could not connect to the database")
    }

    if err := models.MigrateBooks(db); err != nil {
        log.Fatal("Could not migrate database")
    }

    repo := handlers.Repository{DB: db}

    app := fiber.New()
    routes.SetupRoutes(app, &repo)

    log.Fatal(app.Listen(":8080"))
}