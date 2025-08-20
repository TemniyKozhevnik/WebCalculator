package main

import (
	"Calculator_API/internal/calculationService"
	"Calculator_API/internal/handlers"
	"Calculator_API/internal/db"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}