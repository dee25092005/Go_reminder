package main

import (
	"fmt"
	"go-onboarding/api"
	"go-onboarding/storage"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("starting Live Postgres testing...")
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load env: %v\n", err)
		return
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		fmt.Println("no database url provided")
		os.Exit(1)
	}

	store, err := storage.NewSQLStore(connStr)
	if err != nil {
		fmt.Printf("failed to create store: %v\n", err)
		return
	}

	UserHandler := api.NewUserHandler(store)
	fmt.Println("starting server...")
	err = http.ListenAndServe(":8080", UserHandler)
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}

}
