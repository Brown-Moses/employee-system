package repository

import (
	"codewithwuruem/internal/model"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(employee *model.EmployeeDTO) error {
	var existingEmployee model.EmployeeDTO
	if err := r.db.Where("email = ?", employee.Email).First(&existingEmployee).Error; err == nil {
		return errors.New("employee with this email already exists")
	}

	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) GetEmployeeByID(id uint) (*model.EmployeeDTO, error) {
	var employee model.EmployeeDTO
	if err := r.db.Where("employee_id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) UpdateEmployee(employee *model.EmployeeDTO) error {
	return r.db.Save(employee).Error
}

func (r *EmployeeRepository) DeleteEmployee(id uint) error {
	return r.db.Where("employee_id = ?", id).Delete(&model.EmployeeDTO{}).Error
}
