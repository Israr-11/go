package models

// User represents the user table in database
type User struct {
    // ID column, primary key in DB, serialized as "id" in JSON
    ID       uint   `gorm:"primaryKey" json:"id"` 
    // Name column, serialized as "name" in JSON
    Name     string `json:"name"` 
    // Email column, unique in DB, serialized as "email"               
    Email    string `gorm:"unique" json:"email"`
     // Password column, hidden in JSON response  
    Password string `json:"-"`                   
}