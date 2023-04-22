package main

import (
	"bmiz/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading dotenv file")
		}

		testdodotenv := os.Getenv("DOTENV")
		log.Println(testdodotenv)

	}

	DOSERVICE := os.Getenv("DOSYSTEMDSERIVCE")
	log.Println("DO SERVICE:", DOSERVICE)

	routes.HandleRequests()
}
