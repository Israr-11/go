package configuration

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"        // load .env file
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Config holds DB connection info
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

// LoadConfig loads environment variables
func LoadConfig() *Config {
    godotenv.Load() // loads .env automatically
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
    }
}

// ConnectDB connects to PostgreSQL using GORM
func ConnectDB(cfg *Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}