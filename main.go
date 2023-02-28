package main

import (
	"context"
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/onee-only/mempool-manager/api"
	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
)

func main() {

	// .env
	if os.Getenv("DEBUG") == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found")
		}
	}

	// sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://6b453e7cb0a74727af0a5dd75b443ee5@o4504740884185088.ingest.sentry.io/4504757346959360",
		TracesSampleRate: 1.0,
	})
	lib.HandleErr(err)

	// MongoDB
	client := db.InitDatabase()
	defer client.Disconnect(context.TODO())

	// gin
	api.InitServer()
}
