package handler

import (
	"codewithwuruem/internal/db"
	"codewithwuruem/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {

	var employees []model.EmployeeDTO
	if err := db.DB.Preload("Manager").Find(&employees).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func CreateSale(c *gin.Context) {
	var sale model.Sale
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
	var request model.ServiceRequest
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
