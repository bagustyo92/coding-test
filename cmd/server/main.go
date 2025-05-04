package main

import (
	"wallet/internal/handler"
	"wallet/internal/repository"
	"wallet/internal/service"
	"wallet/routes"
)

func main() {
	userRepo := repository.NewUserRepository()
	disburseService := service.NewDisbursementService(userRepo)
	disburseHandler := handler.NewDisbursementHandler(disburseService)

	router := routes.SetupRouter(disburseHandler)
	router.Run(":8080")
}
