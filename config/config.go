package config

import (
	"log"
	"os"
)

var port string

func LoadConfig() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT environment variable not set. Using default port 8080.")
	}
}

func GetPort() string {
	return port
}
