package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config armazena as configurações da aplicação
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig contém configurações do servidor
type ServerConfig struct {
	Port int
	Env  string
}

// DatabaseConfig contém configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Load carrega as configurações das variáveis de ambiente
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 8000),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "vilar"),
			Password: getEnv("DB_PASSWORD", "vilar123"),
			DBName:   getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

// GetDSN retorna a string de conexão do banco de dados
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Database.Host,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.Port,
		c.Database.SSLMode,
	)
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtém uma variável de ambiente como int ou retorna um valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Erro ao converter %s para int, usando valor padrão: %d", key, defaultValue)
		return defaultValue
	}
	return value
}
