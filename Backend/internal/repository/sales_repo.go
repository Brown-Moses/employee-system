package repository

import (
	"codewithwuruem/internal/model"
	"errors"

	"gorm.io/gorm"
)

type SalesRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{db: db}
}

// only sales rep can handle sales
func (r *SalesRepository) CreateSale(sale *model.SaleDTO) error {
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
func (r *SalesRepository) CreateServiceRequest(request *model.ServiceRequestDTO) error {
	var employee model.EmployeeDTO
	if err := r.db.First(&employee, request.TechnicianID).Error; err != nil {
		return errors.New("technician not found")
	}

	if employee.JobTitle != "Technician" {
		return errors.New("Only technicians can be assigned to service requests")
	}

	return r.db.Create(request).Error
}
