package model

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID  uint       `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	FirstName   string     `gorm:"column:first_name" json:"first_name"`
	LastName    string     `gorm:"column:last_name" json:"last_name"`
	DateOfBirth *time.Time `gorm:"column:date_of_birth" json:"date_of_birth,omitempty"`
	Gender      Gender     `gorm:"column:gender" json:"gender"`
	Phone       string     `gorm:"column:phone" json:"phone"`
	Email       string     `gorm:"column:email;unique" json:"email"`
	HireDate    *time.Time `gorm:"column:hire_date" json:"hire_date,omitempty"`
	Salary      float64    `gorm:"column:salary" json:"salary"`
	JobTitle    JobTitle   `gorm:"column:job_title" json:"job_title"`
	ManagerID   *uint      `gorm:"column:manager_id" json:"manager_id,omitempty"`
	Manager     *Employee  `gorm:"foreignKey:ManagerID;references:EmployeeID" json:"manager,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Manager struct {
	EmployeeID      uint      `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	ManagementLevel string    `gorm:"column:management_level" json:"management_level"`
	Department      string    `gorm:"column:department" json:"department"`
	Employee        *Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID;constraint:OnDelete:CASCADE" json:"employee,omitempty"`
}

type SalesRepresentative struct {
	EmployeeID     uint      `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	CommissionRate float64   `gorm:"column:commission_rate" json:"commission_rate"`
	SalesTarget    float64   `gorm:"column:sales_target" json:"sales_target"`
	Employee       *Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID;constraint:OnDelete:CASCADE" json:"employee,omitempty"`
}

type Technician struct {
	EmployeeID          uint      `gorm:"primaryKey;column:employee_id" json:"employee_id"`
	SkillSpecialization string    `gorm:"column:skill_specialization" json:"skill_specialization"`
	CertificationLevel  string    `gorm:"column:certification_level" json:"certification_level"`
	Employee            *Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID;constraint:OnDelete:CASCADE" json:"employee,omitempty"`
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
	SaleID        uint                 `gorm:"primaryKey;column:sale_id" json:"sale_id"`
	SaleDate      time.Time            `gorm:"column:sale_date;not null" json:"sale_date"`
	Amount        float64              `gorm:"column:amount;not null" json:"amount"`
	PaymentMethod string               `gorm:"column:payment_method" json:"payment_method"`
	CustomerName  string               `gorm:"column:customer_name" json:"customer_name"`
	EmployeeID    *uint                `gorm:"column:employee_id" json:"employee_id,omitempty"`
	SalesRep      *SalesRepresentative `gorm:"foreignKey:EmployeeID;references:EmployeeID;constraint:OnDelete:SET NULL" json:"sales_rep,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ServiceRequest struct {
	RequestID      uint        `gorm:"primaryKey;column:request_id" json:"request_id"`
	RequestDate    time.Time   `gorm:"column:request_date;not null" json:"request_date"`
	Description    string      `gorm:"column:description;type:text" json:"description"`
	Status         string      `gorm:"column:status" json:"status"`
	Priority       string      `gorm:"column:priority" json:"priority"`
	CompletionDate *time.Time  `gorm:"column:completion_date" json:"completion_date,omitempty"`
	TechnicianID   *uint       `gorm:"column:technician_id" json:"technician_id,omitempty"`
	Technician     *Technician `gorm:"foreignKey:TechnicianID;references:EmployeeID;constraint:OnDelete:SET NULL" json:"technician,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
