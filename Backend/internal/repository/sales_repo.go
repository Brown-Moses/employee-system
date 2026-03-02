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
	var salesRep model.SalesRepresentative
	if err := r.db.First(&salesRep, sale.EmployeeID).Error; err != nil {
		return errors.New("sales representative not found")
	}

	return r.db.Create(sale).Error
}

// only technicians can handle service requests
func (r *SalesRepository) CreateServiceRequest(request *model.ServiceRequestDTO) error {
	var technician model.Technician
	if err := r.db.First(&technician, request.TechnicianID).Error; err != nil {
		return errors.New("technician not found")
	}

	return r.db.Create(request).Error
}
