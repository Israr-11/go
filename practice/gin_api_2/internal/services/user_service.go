package services

/*
WHAT THIS FILE DOES (Simple Explanation)

This file contains the BUSINESS LOGIC for users.

Flow of the application:

HTTP Request
   ↓
Controller/Handler
   ↓
UserService (THIS FILE) → applies business logic
   ↓
UserRepository → talks to database
   ↓
Database

Example when creating a user:

Request → CreateUser(name, email, password)
        ↓
Password is hashed (security step)
        ↓
User struct is created
        ↓
Repository is called to save user in database

Think of this like the "service layer" in Node.js or Flask apps.
*/

import (
    "gin-quickstart/internal/models"
    "gin-quickstart/internal/repository"
    "gin-quickstart/pkg/utils"
)

// UserService contains business logic for users
type UserService struct {


// *Type -> Pointer (reference to object)

// Example:  repo *UserRepository

// Think: store reference to repository

// Like JS object reference.

    // repo holds the repository that talks to the database
    // *repository.UserRepository means we store a reference to the repository
    // (in simple terms: we keep the same repo object instead of copying it)
    repo *repository.UserRepository
}

// Constructor
// Creates a new UserService and attaches the repository to it
func NewUserService(repo *repository.UserRepository) *UserService {

    // &UserService means "create a new UserService and return its reference"
    return &UserService{repo: repo}
}

// CreateUser hashes password and calls repository
func (s *UserService) CreateUser(name, email, password string) error {

    // (s *UserService)
    // s is the current service instance
    // similar to "this" in other languages

    // Hash the password before storing in DB
    // utils.HashPassword probably uses bcrypt internally
    hashedPassword, err := utils.HashPassword(password)

	//We check whether the err variable is nil, if nil operation succeeded and
	// if not then it's mean something went wrong.

    if err != nil {
        return err
    }

    // Create a new User struct

	// Same as=> var user struct = &models.User
    user := &models.User{

        // &models.User means we create the user and return its reference
        // simple thinking: create the user object

        Name:     name,
        Email:    email,
        Password: hashedPassword,
    }

    // Call repository to insert user into database
    return s.repo.Create(user)
}

// GetAllUsers calls repository to get all users
func (s *UserService) GetAllUsers() ([]models.User, error) {

    // Simply ask the repository to fetch users from DB
    return s.repo.GetAll()
}