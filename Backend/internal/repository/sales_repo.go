package repository

import (
	"codewithwuruem/internal/model"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// getting employee,(Manager)
func (r *Repository) GetEmployeesByID() ([]model.Employee, error) {
	var employees []model.Employee
	if err := r.db.Preload("Manager").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// only sales rep can handle sales
func (r *Repository) CreateSale(sale *model.SaleDTO) error {
	var employee model.EmployeeDTO
	if err := r.db.First(&employee, sale.SaleRepresentativeID).Error; err != nil {
		return errors.New("sale representative not found")
	}

	if employee.JobTitle != "Sales Representative" {
		return errors.New("Only sales representative can create sales")
	}

	return r.db.Create(sale).Error
}

// only technicians can handle service requests
func (r *Repository) CreateServiceRequest(request *model.ServiceRequestDTO) error {
	var employee model.EmployeeDTO
	if err := r.db.First(&employee, request.TechnicianID).Error; err != nil {
		return errors.New("technician not found")
	}

	if employee.JobTitle != "Technician" {
		return errors.New("Only technicians can be assigned to service requests")
	}

	return r.db.Create(request).Error
}
