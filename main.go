package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/onee-only/mempool-manager/api"
	"github.com/onee-only/mempool-manager/db"
)

func main() {

	// .env
	if os.Getenv("DEBUG") == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found")
		}
	}

	// gin
	api.InitServer()

	// MongoDB
	client := db.InitDatabase()
	defer client.Disconnect(context.TODO())
}
