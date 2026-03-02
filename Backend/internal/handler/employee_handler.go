package handler

import (
	"codewithwuruem/internal/db"
	"codewithwuruem/internal/model"
	"codewithwuruem/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEmployees(c *gin.Context) {
	var employee model.EmployeeDTO
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	empRepo := repository.NewEmployeeRepository(db.DB)
	if err := empRepo.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func GetEmployees(c *gin.Context) {
	var employees []model.Employee
	if err := db.DB.Preload("Manager").Find(&employees).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	if err := db.DB.Preload("Manager").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	if err := db.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updatedData model.EmployeeDTO
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if employee.ManagerID != nil && updatedData.ManagerID != nil && *employee.ManagerID != *updatedData.ManagerID {
		c.JSON(400, gin.H{"error": "Manager cant be their own manager"})
	}

	c.JSON(http.StatusCreated, employee)
}

func DeleteEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&model.Employee{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
