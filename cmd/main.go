package main

import (
	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/database"
	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/log", services.PostLogService)

	r.Run()
}
