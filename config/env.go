package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAppPort() string {
	return getEnv("APP_PORT", "8080")
}

func BuildPostgresDSN() (string, error) {
	host := firstNonEmpty(os.Getenv("DB_HOST"), os.Getenv("POSTGRES_HOST"), "postgres")
	port := firstNonEmpty(os.Getenv("DB_PORT"), "5432")
	user := firstNonEmpty(os.Getenv("DB_USER"), os.Getenv("POSTGRES_USER"), "root")
	password := firstNonEmpty(os.Getenv("DB_PASSWORD"), os.Getenv("POSTGRES_PASSWORD"), "root")
	name := firstNonEmpty(os.Getenv("DB_NAME"), os.Getenv("POSTGRES_DB"), "lab-ci-cd-go")
	sslMode := firstNonEmpty(os.Getenv("DB_SSLMODE"), "disable")
	appEnv := strings.ToLower(firstNonEmpty(os.Getenv("APP_ENV"), "development"))

	if _, err := strconv.Atoi(port); err != nil {
		return "", fmt.Errorf("DB_PORT invalido: %w", err)
	}

	if appEnv == "production" {
		if os.Getenv("DB_USER") == "" && os.Getenv("POSTGRES_USER") == "" {
			return "", fmt.Errorf("DB_USER/POSTGRES_USER obrigatorio em producao")
		}
		if os.Getenv("DB_PASSWORD") == "" && os.Getenv("POSTGRES_PASSWORD") == "" {
			return "", fmt.Errorf("DB_PASSWORD/POSTGRES_PASSWORD obrigatorio em producao")
		}
		if os.Getenv("DB_NAME") == "" && os.Getenv("POSTGRES_DB") == "" {
			return "", fmt.Errorf("DB_NAME/POSTGRES_DB obrigatorio em producao")
		}
		if sslMode == "disable" {
			return "", fmt.Errorf("DB_SSLMODE nao pode ser disable em producao")
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, name, port, sslMode,
	)

	return dsn, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}

	return ""
}
