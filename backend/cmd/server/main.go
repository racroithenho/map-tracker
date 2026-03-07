package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/racroithenho/map-tracker/backend/internal/handler"
	"github.com/racroithenho/map-tracker/backend/internal/repository"
	routes "github.com/racroithenho/map-tracker/backend/internal/router"
	"github.com/racroithenho/map-tracker/backend/internal/service"
	"github.com/racroithenho/map-tracker/backend/migrations"
)

func main() {
	db, err := migrations.Connect()
	if err != nil {
		log.Fatalf("Failed to connect PostgreSQL: %v", err)
	}

	deviceRepo := repository.NewDeviceRepository(db)
	deviceService := service.NewDeviceService(deviceRepo)
	deviceHandler := handler.NewDeviceHandler(deviceService)

	r := gin.Default()

	routes.RegisterDeviceRoutes(r, deviceHandler)

	// Start server on port 8080 (default)
 	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}