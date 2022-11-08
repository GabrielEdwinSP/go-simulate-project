package main

import (
	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/database"
	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/domain"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&domain.Request{}, &domain.Response{})
}
