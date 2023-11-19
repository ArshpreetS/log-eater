package main

import (
	"fmt"
	"os"

	"github.com/ArshpreetS/log-ingester/database"
	"github.com/ArshpreetS/log-ingester/routes"
	"github.com/joho/godotenv"
)

func main() {

	// Loading the env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error with loading the env file")
	}

	// Setting the elastic database client
	es_client, err := database.CreateClient()
	if err != nil {
		fmt.Println("There was an error with elasticsearch client:", err)
	}

	r := routes.SetupRoutes(es_client)
	r.Run(":" + os.Getenv("PORT"))
}
