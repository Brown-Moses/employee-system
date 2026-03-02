package main

import (
	"codewithwuruem/internal/db"
	"codewithwuruem/internal/handler"
	"codewithwuruem/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {

	err := db.InitDB()
	if err != nil {
		panic(err)
	} else {
		println("connection successful!")
	}

	db.DB.AutoMigrate(
		&model.Employee{},
		&model.Manager{},
		&model.SalesRepresentative{},
		&model.Technician{},
		&model.Sale{},
		&model.ServiceRequest{},
		&model.Customer{},
	)

	r := gin.Default()

	//getting ready!!
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Everything is working fine!",
		})
	})

	r.POST("/employees", handler.AddEmployees)
	r.GET("/employees/:id", handler.GetEmployeeByID)
	r.PUT("/employees/:id", handler.UpdateEmployee)
	r.DELETE("/employees/:id", handler.DeleteEmployeeByID)
	r.GET("/employees", handler.GetEmployees)
	r.POST("/sales", handler.CreateSale)
	r.POST("/service-requests", handler.CreateServiceRequest)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
