package models

// User represents the user table in database
type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"-"` // "-" hides password in JSON response
}