package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	UserPoolClientId   string
	AwsRegion          string
	TokenExpirySeconds int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found or failed to load: %v\n", err)
	}

	port := getEnv("PORT", "3000")
	userPoolClientId := getEnv("USER_POOL_CLIENT_ID", "")
	awsRegion := getEnv("AWS_REGION", "us-east-1")
	tokenExpirySeconds := getEnvAsInt("TOKEN_EXPIRY_SECONDS", 3600)

	config := &Config{
		Port:               port,
		UserPoolClientId:   userPoolClientId,
		AwsRegion:          awsRegion,
		TokenExpirySeconds: tokenExpirySeconds,
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
