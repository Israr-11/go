// This file loads database credentials from a .env file and creates a 
// PostgreSQL connection using GORM so the rest of the application can interact 
// with the database.
package configuration

import (
    "fmt" // for formatting strings
    "os"  // for reading environment variables

    "github.com/joho/godotenv" // loads .env file into environment variables
    "gorm.io/driver/postgres"   // PostgreSQL driver for GORM
    "gorm.io/gorm"              // GORM ORM
)

// Config holds DB connection info. We create one instance of this struct
// after reading values from environment variables.
type Config struct {
    DBHost     string // database host, e.g., "localhost"
    DBPort     string // database port, e.g., "5432"
    DBUser     string // database username
    DBPassword string // database password
    DBName     string // database name
}

// LoadConfig loads environment variables from the .env file
// and returns a pointer to a Config struct containing all DB settings.
func LoadConfig() *Config {
    // Load the .env file into environment variables
    // If .env is missing, it silently continues (you may want to handle errors 
    // in prod)
    godotenv.Load()
    if err := godotenv.Load(); err != nil {
    fmt.Println(".env file not found")
}

    // Create a Config object by reading the environment variables
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
    }
}

// ConnectDB connects to PostgreSQL using GORM and returns the *gorm.DB object.
// This object is used throughout the app to interact with the database.
func ConnectDB(cfg *Config) (*gorm.DB, error) {
    // Build the DSN (Data Source Name) string for PostgreSQL
    // fmt.Sprintf formats a string by replacing %s with the arguments in order.
    // Example: "host=localhost user=postgres password=pass dbname=mydb 
    // port=5432 sslmode=disable"
    
    // Raw Go equivalent using concatenation would be:
    // dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + 
    // cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort 
    // + " sslmode=disable"
    
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
    )
    // Explanation: fmt.Sprintf takes the format string (first argument)
    // and replaces each %s with the corresponding argument in order.
    // So the above is exactly like building the string manually using '+' in Go.
    
    // Open a connection to PostgreSQL using GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        // Return nil and the error if connection fails
        return nil, err
    }

    // Return the database connection object
    return db, nil
}