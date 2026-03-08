/*
WHAT THIS FILE DOES (Simple Explanation)

This file connects HTTP URLs to handler functions.

Flow:

HTTP Request
    ↓
Router (THIS FILE decides which handler runs)
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
Database

Example:

POST /api/users
    ↓
userHandler.CreateUser()

GET /api/users
    ↓
userHandler.GetUsers()
*/


package routes

import (
    "github.com/gin-gonic/gin"
    "gin-quickstart/internal/handlers"
)

// SetupRouter wires endpoints to handlers
func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {

    // gin.Default() creates a Gin router with:
    // - logger middleware
    // - recovery middleware (prevents crashes)
    r := gin.Default()

    // Group creates a route prefix
    // All routes inside will start with /api
    api := r.Group("/api")

    {
        // POST /api/users
        // When this endpoint is hit → call CreateUser handler
        api.POST("/users", userHandler.CreateUser)

        // GET /api/users
        // When this endpoint is hit → call GetUsers handler
        api.GET("/users", userHandler.GetUsers)
    }

    // Return the router so main.go can start the server
    return r
}