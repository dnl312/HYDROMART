package main

import (
	"log"
	"merchant/config"
	"merchant/controller"
	"merchant/repo"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	config.ClearPreparedStatements()
	log.Printf("Connected to database")
	defer config.CloseDB()
	bookRepository := repo.NewMerchantRepository(db)
	bookController := controller.NewMerchantController(&bookRepository)

	config.ListenAndServeGrpc(&bookController)
}