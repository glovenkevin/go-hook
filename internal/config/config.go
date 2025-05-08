package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	// Try to read from .env file
	loadEnvFromFile(".env")

	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

// loadEnvFromFile reads a file line by line and sets environment variables
func loadEnvFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// Continue even if .env file is not found
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments and empty lines
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set environment variable
		os.Setenv(key, value)
	}
}

// getEnv returns the value of an environment variable or a default value if not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
