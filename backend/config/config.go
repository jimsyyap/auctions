// config/config.go
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	// Get the current file's directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	
	// Load .env file from the root directory
	err := godotenv.Load(filepath.Join(dir, "..", ".env"))
	if err != nil {
		log.Printf("Warning: .env file not found or cannot be loaded: %v", err)
		// Continue execution as environment variables might be set elsewhere (e.g., in production)
	}
}

// DBConfig holds database connection parameters
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetDBConfig returns database configuration from environment variables
func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "auctions_db"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}
}

// GetDBConnectionString returns formatted DSN string for database connection
func GetDBConnectionString() string {
	config := GetDBConfig()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// GetServerConfig returns server configuration from environment variables
func GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
	}
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
