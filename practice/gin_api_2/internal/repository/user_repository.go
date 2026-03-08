/*
QUICK POINTER & STRUCT CHEAT SHEET:

1️⃣ * (pointer type)
   - Example: db *gorm.DB
   - Means “db is a pointer to a gorm.DB object”
   - Pointer = stores memory address of the object instead of copying it
   - Saves memory and allows shared modification

2️⃣ & (address of)
   - Example: &users
   - Means “give me the memory address of the variable users”
   - Needed when a function should modify the original variable

3️⃣ (r *UserRepository)
   - r = instance of UserRepository
   - *UserRepository = pointer receiver, allows method to access & modify same 
   struct
   - Think of r as “this” in other languages

4️⃣ Methods & Constructors
   - func (r *UserRepository) Create(...) {...} → Create is a 
   method for the struct
   - func NewUserRepository(db *gorm.DB) *UserRepository → creates new 
   instance of repository

This file: handles DB operations for the User model using GORM, keeps
service layer clean, and only deals with database logic.
*/

/*
1. Declaration: You declare a pointer variable using an asterisk (*) before its 
data type. For example, var ptr *int declares a pointer to an integer. 
This means that the ptr variable will hold the memory address of an integer.


2. Initialization: To make a pointer hold the memory address of a variable, you 
use the address-of operator (&). For example, ptr = &num assigns the memory 
address of variable num to the pointer ptr.

3. Dereferencing: To access the value stored at the memory location pointed to 
by a pointer, we use the dereference operator (*). For example, *ptr retrieves 
the value stored at the address ptr points to.
*/

package repository

import (
    "gin-quickstart/internal/models" // import the User model struct

    "gorm.io/gorm" // GORM database ORM
)

// UserRepository handles DB operations.
// This struct acts like a wrapper around the database connection.
type UserRepository struct {
    db *gorm.DB // pointer to the database connection object
}

// Constructor function that creates a new UserRepository.
// It receives the database connection and stores it in the struct.
func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
    // &UserRepository{} creates a pointer to a new struct instance.
    // db: db means "assign the passed db parameter to the struct field db".
}

// Create inserts a new user into DB
func (r *UserRepository) Create(user *models.User) error {

	//r points to the UserRepository
    // (r *UserRepository) means this function is a METHOD
    // attached to the UserRepository struct.
    // r is the instance of the repository (like "this" in other languages).

    // user *models.User means we pass a pointer to a User struct.
    // Using a pointer avoids copying the entire struct and lets GORM modify it.

    return r.db.Create(user).Error

    // r.db -> the database connection stored inside the repository
    // Create(user) -> GORM inserts the user into the database
    // .Error -> returns the error from the operation
}

// GetAll fetches all users from DB
func (r *UserRepository) GetAll() ([]models.User, error) {

    // []models.User means a slice (**** dynamic array ****) of User structs
    var users []models.User

    // Find(&users) tells GORM to fetch all rows from the users table
    // and populate the users slice.

    err := r.db.Find(&users).Error

    // &users means we pass the memory address so GORM can fill it.

    return users, err
}