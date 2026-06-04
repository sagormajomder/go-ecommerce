package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
	DB          DBConfig
}

var cnf *Config

func loadConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("Service Name is required")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("Http Port is required")
	}

	port, err := strconv.Atoi(httpPort)

	if err != nil {
		log.Fatal("port must be in number:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Jwt secret is required")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB Host is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB Port is required")
	}

	dbPrt, err := strconv.Atoi(dbPort)

	if err != nil {
		log.Fatal("port must be in number:", err)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB Name is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB User is required")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		log.Fatal("DB Pass is required")
	}

	sslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enableSSLMode, err := strconv.ParseBool(sslMode)

	if err != nil {
		log.Fatal("Invalid enable ssl mode value", err)
	}

	dbConfig := DBConfig{
		Host:          host,
		Port:          dbPrt,
		Name:          dbName,
		User:          dbUser,
		Password:      dbPass,
		EnableSSLMode: enableSSLMode,
	}

	cnf = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		JWTSecret:   jwtSecret,
		DB:          dbConfig,
	}
}

func GetConfig() *Config {
	if cnf == nil {
		loadConfig()
	}
	return cnf
}
