package model

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID uint      `gorm:"primaryKey" json:"employee_id"`
	FirstName  string    `gorm:"column:first_name" json:"first_name"`
	LastName   string    `gorm:"column:last_name" json:"last_name"`
	JobTitle   JobTitle  `gorm:"column:job_title" json:"job_title"`
	Gender     Gender    `gorm:"column:gender" json:"gender"`
	Email      string    `gorm:"column:email;unique" json:"email"`
	Salary     float64   `gorm:"column:salary" json:"salary"`
	ManagerID  *uint     `gorm:"column:manager_id" json:"manager_id"`
	Manager    *Employee `gorm:"foreignKey:ManagerID;references:EmployeeID" json:"manager,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Customer struct {
	CustomerID uint   `gorm:"primaryKey" json:"customer_id"`
	FirstName  string `gorm:"column:first_name" json:"first_name"`
	LastName   string `gorm:"column:last_name" json:"last_name"`
	Email      string `gorm:"column:email" json:"email"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Sale struct {
	SaleID               uint     `gorm:"primaryKey" json:"sale_id"`
	SaleDate             string   `gorm:"column:sale_date" json:"sale_date"`
	Amount               float64  `gorm:"column:amount" json:"amount"`
	SaleRepresentativeID uint     `gorm:"not null" json:"sale_representative_id"`
	SaleRepresentative   Employee `gorm:"foreignKey:SaleRepresentativeID;references:EmployeeID" json:"sale_representative"`
	CustomerID           uint     `gorm:"not null" json:"customer_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ServiceRequest struct {
	RequestID    uint     `gorm:"primaryKey" json:"request_id"`
	Description  string   `gorm:"column:description" json:"description"`
	TechnicianID uint     `gorm:"not null" json:"technician_id"`
	Technician   Employee `gorm:"foreignKey:TechnicianID;references:EmployeeID" json:"technician"`
	CustomerID   uint     `gorm:"not null" json:"customer_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
