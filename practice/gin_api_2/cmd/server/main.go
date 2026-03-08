package main

/*
WHAT THIS FILE DOES (Simple Explanation)

This is the ENTRY POINT of the application.

main() runs when the program starts.

This file connects everything together:

1. Load configuration (.env)
2. Connect to database
3. Run database migrations
4. Create repository layer
5. Create service layer
6. Create handler layer
7. Setup routes
8. Start HTTP server

Think of this like the "app.js" or "server.js" in Node.js.
*/

import (
    "log"

    "gin-quickstart/configuration"
    "gin-quickstart/internal/handlers"
    "gin-quickstart/internal/models"
    "gin-quickstart/internal/repository"
    "gin-quickstart/internal/routes"
    "gin-quickstart/internal/services"
)

func main() {

    // STEP 1: Load environment variables from .env
    cfg := configuration.LoadConfig()

    // STEP 2: Connect to PostgreSQL database
    db, err := configuration.ConnectDB(cfg)
    if err != nil {

        // log.Fatalf prints error and stops program
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    // STEP 3: Auto-migrate database schema
    // This creates/updates the users table automatically
    db.AutoMigrate(&models.User{})

    // STEP 4: Initialize repository layer
    // Repository talks to database
    userRepo := repository.NewUserRepository(db)

    // STEP 5: Initialize service layer
    // Service contains business logic
    userService := services.NewUserService(userRepo)

    // STEP 6: Initialize handler layer
    // Handler deals with HTTP requests
    userHandler := handlers.NewUserHandler(userService)

    // STEP 7: Setup router (connect endpoints to handlers)
    r := routes.SetupRouter(userHandler)

    // STEP 8: Start HTTP server
    // Server will run on http://localhost:8080
    r.Run(":8080")
}