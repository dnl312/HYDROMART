package main

// import (
// 	"auth/config"
// 	"auth/controller"
// 	"auth/repo"
// 	"log"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	db, err := config.InitDB()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	config.ClearPreparedStatements()

// 	defer config.CloseDB()
// 	userRepository := repo.NewUserRepository(db)
// 	userController := controller.NewAuthController(userRepository)

// 	config.ListenAndServeGrpc(&userController)
// }


