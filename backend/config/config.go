package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var cnf Config

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

	cnf = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}
}

func GetConfig() Config {
	loadConfig()
	return cnf
}
