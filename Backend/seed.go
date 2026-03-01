package main

import (
	"log"

	"codewithwuruem/internal/db"
	"codewithwuruem/internal/model"
)

func run() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	} else {
		println("connection successful!")
	}

	// drop tables to ensure schema is recreated cleanly (development only)
	_ = db.DB.Migrator().DropTable(&model.Employee{}, &model.Customer{}, &model.Sale{}, &model.ServiceRequest{})
	if err := db.DB.AutoMigrate(&model.Employee{}, &model.Sale{}, &model.ServiceRequest{}, &model.Customer{}); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}

	// seed employees
	employees := []model.Employee{
		{EmployeeID: 1, FirstName: "Moses", LastName: "Justin", JobTitle: model.JobTech, Gender: model.GenderFemale, Salary: 75000},
		{EmployeeID: 2, FirstName: "Peter", LastName: "Jones", JobTitle: model.JobManager, Gender: model.GenderMale, Salary: 90000, ManagerID: ptrUint(1)},
		{EmployeeID: 3, FirstName: "Sarah", LastName: "Williams", JobTitle: model.JobSales, Gender: model.GenderFemale, Salary: 65000, ManagerID: ptrUint(2)},
	}
	for _, e := range employees {
		if err := db.DB.Create(&e).Error; err != nil {
			log.Printf("failed to insert employee %v: %v", e.EmployeeID, err)
		}
	}

	customers := []model.Customer{
		{CustomerID: 1, FirstName: "Justin", LastName: "Brown", Email: "brown@example.com"},
		{CustomerID: 2, FirstName: "Eve", LastName: "Davis", Email: "eve@example.com"},
	}
	for _, c := range customers {
		if err := db.DB.Create(&c).Error; err != nil {
			log.Printf("failed to insert customer %v: %v", c.CustomerID, err)
		}
	}

	log.Println("seeding complete")
}

func ptrUint(u uint) *uint { return &u }
