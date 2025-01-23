package main

import (
	"log"
	"merchant/config"
	"merchant/controller"
	"merchant/repo"
	"merchant/service"

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


	conn, mbChan := config.InitMessageBroker()
	defer conn.Close()
	
	messageBrokerService := service.NewMessageBroker(mbChan)

	config.ClearPreparedStatements()
	log.Printf("Connected to database")
	defer config.CloseDB()
	merchantRepository := repo.NewMerchantRepository(db)
	orderRepository := repo.NewOrderRepository(db)
	merchantController := controller.NewMerchantController(&merchantRepository, &orderRepository, messageBrokerService)

	config.ListenAndServeGrpc(&merchantController)
}
