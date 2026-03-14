package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/joho/godotenv"

    "Synconomics/config"
    "Synconomics/handlers"
    "Synconomics/repositories"
    "Synconomics/services"
)

func main() {
    godotenv.Load()

    // connect database
    config.ConnectDB()

    // dependency injection
    userRepo    := repositories.NewUserRepository(config.DB)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    app := fiber.New()
    app.Use(recover.New())
    app.Use(logger.New())

    // public routes
    api := app.Group("/api")
    api.Post("/register", userHandler.Register)
    api.Post("/login",    userHandler.Login)

    // protected routes
    api.Get("/profile", handlers.JWTMiddleware, userHandler.Profile)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(app.Listen(":" + port))
}