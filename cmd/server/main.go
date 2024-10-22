package main

import (
	"github.com/madhwan-codes/authify/internal/config"
	"log"
)

func main() {
	_, loadConfigErr := config.LoadConfig()
	if loadConfigErr != nil {
		log.Fatal("Failed to load configurations, error: ", loadConfigErr)
	}

}
