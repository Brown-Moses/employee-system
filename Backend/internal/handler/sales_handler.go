package handler

import (
	"codewithwuruem/internal/db"
	"codewithwuruem/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSale(c *gin.Context) {

	var sale model.SaleDTO
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.DB.Create(&sale).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create sale"})
		return
	}
	c.JSON(http.StatusCreated, sale)

}

func CreateServiceRequest(c *gin.Context) {
	var request model.ServiceRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.DB.Create(&request).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create service request"})
		return
	}
	c.JSON(http.StatusCreated, request)
}
