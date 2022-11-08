package main

import (
	"github.com/GabrielEdwinSP/go-simutale-pajak.git/internal/database"
	"github.com/GabrielEdwinSP/go-simutale-pajak.git/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	db := database.Database()
	logService := services.NewLogService(db)

	route.POST("/log", logService.PostLogService)
	route.Run(":3306")
}
