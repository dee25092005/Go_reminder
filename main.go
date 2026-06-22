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
	loginHandler := api.NewLoginHandler()

	protectedUserChain := api.Logger(api.RequireAuth(UserHandler))

	publicLoginChain := api.Logger(loginHandler)

	mux := http.NewServeMux()
	mux.Handle("/users", protectedUserChain)
	mux.Handle("/login", publicLoginChain)

	fmt.Println("starting server...")
	fmt.Println("starting server on http://localhost:8080 ...")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}

}
