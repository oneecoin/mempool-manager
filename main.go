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

	// gin
	api.InitServer()

	// .env
	if os.Getenv("DEBUG") == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found")
		}
	}

	// MongoDB
	client := db.InitDatabase()
	defer client.Disconnect(context.TODO())
}
