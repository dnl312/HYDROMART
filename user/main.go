package main

import (
	"log"
	"user/config"
	"user/controller"
	"user/repo"

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

	userRepository := repo.NewUserRepository(db)
	userController := controller.NewMerchantController(&userRepository)

	config.ListenAndServeGrpc(&userController)
}
