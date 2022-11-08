package services

import (
	"net/http"

	"github.com/GabrielEdwinSP/go-simulate-project.git/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogService struct {
	db *gorm.DB
}

func NewLogService(database *gorm.DB) *LogService {
	return &LogService{
		db: database,
	}
}

func (es LogService) PostLogService(c *gin.Context) {
	var request domain.Request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}
	var response domain.Response
	c.SecureJSON(
		http.StatusOK, gin.H{
			"totalgross": response.TotalGross,
		})
	c.JSON(200, response)
}
