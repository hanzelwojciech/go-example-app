package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Initialize() string {
	mode := strings.Trim(os.Getenv("MODE"), " ")

	if mode == "local" {
		fmt.Println("Importing custom env file")

		err := godotenv.Load("../.env")

		if err != nil {
			log.Fatalln("[dotenv] Error loading .env file", err.Error())
		}
	}

	return mode
}
