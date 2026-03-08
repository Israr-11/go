/*
WHAT THIS FILE DOES (Simple Explanation)

This file defines the HTTP endpoints for users using the Gin web framework.

Flow of a request:

Client (Postman / Browser)
        ↓
HTTP Endpoint (Handler in this file)
        ↓
UserService (business logic)
        ↓
UserRepository (database operations)
        ↓
Database

Example request:

POST /users
{
  "name": "Alice",
  "email": "alice@email.com",
  "password": "123456"
}

Handler will:
1. Read JSON from request
2. Validate it
3. Call UserService
4. Return JSON response
*/

package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gin-quickstart/internal/services"
)

// UserHandler = controller layer (similar to Express or Flask controllers)
type UserHandler struct {

    // service connects this handler to the business logic
    // *services.UserService means we store a reference to the service
    service *services.UserService
}

// Constructor
// Creates a new UserHandler and attaches the service
func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service: service}
}

// Request payload structure
// This struct defines what JSON we expect from the client
type CreateUserRequest struct {

    // `json:"name"` → maps JSON field "name" to this struct field
    // `binding:"required"` → Gin validation (field must exist)
    Name string `json:"name" binding:"required"`

    // email must exist AND must be valid email format
    Email string `json:"email" binding:"required,email"`

    Password string `json:"password" binding:"required"`
}

// CreateUser endpoint
func (h *UserHandler) CreateUser(c *gin.Context) {

    // (h *UserHandler)
    // h is the handler instance (similar to "this" in other languages)

    // c *gin.Context
    // Context contains everything about the HTTP request and response

    var req CreateUserRequest

    // ShouldBindJSON reads JSON body and fills the struct
    // Example request body:
    // { "name":"Alice", "email":"a@email.com", "password":"123" }

    if err := c.ShouldBindJSON(&req); err != nil {

        // If JSON invalid or missing fields → return 400
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Call service layer to create user
    err := h.service.CreateUser(req.Name, req.Email, req.Password)

    if err != nil {

        // If something failed in service or DB
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Success response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
    })
}

// GetUsers endpoint
func (h *UserHandler) GetUsers(c *gin.Context) {

    // Ask service layer for all users
    users, err := h.service.GetAllUsers()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Return users as JSON
    c.JSON(http.StatusOK, users)
}