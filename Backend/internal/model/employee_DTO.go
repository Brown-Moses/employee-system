package model

import "time"

type EmployeeDTO struct {
	EmployeeID  uint       `json:"employee_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Gender      Gender     `json:"gender"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	HireDate    *time.Time `json:"hire_date,omitempty"`
	Salary      float64    `json:"salary"`
	JobTitle    JobTitle   `json:"job_title"`
	ManagerID   *uint      `json:"manager_id,omitempty"`
}

type ManagerDTO struct {
	EmployeeID      uint         `json:"employee_id"`
	ManagementLevel string       `json:"management_level"`
	Department      string       `json:"department"`
	Employee        *EmployeeDTO `json:"employee,omitempty"`
}

type SalesRepresentativeDTO struct {
	EmployeeID     uint         `json:"employee_id"`
	CommissionRate float64      `json:"commission_rate"`
	SalesTarget    float64      `json:"sales_target"`
	Employee       *EmployeeDTO `json:"employee,omitempty"`
}

type TechnicianDTO struct {
	EmployeeID          uint         `json:"employee_id"`
	SkillSpecialization string       `json:"skill_specialization"`
	CertificationLevel  string       `json:"certification_level"`
	Employee            *EmployeeDTO `json:"employee,omitempty"`
}

// CustomerDTO mirrors the Customer model for API requests/responses.
type CustomerDTO struct {
	CustomerID uint   `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}

// SaleDTO represents a sale along with related entities.
type SaleDTO struct {
	SaleID        uint                    `json:"sale_id"`
	SaleDate      time.Time               `json:"sale_date"`
	Amount        float64                 `json:"amount"`
	PaymentMethod string                  `json:"payment_method"`
	CustomerName  string                  `json:"customer_name"`
	EmployeeID    *uint                   `json:"employee_id,omitempty"`
	SalesRep      *SalesRepresentativeDTO `json:"sales_rep,omitempty"`
}

// ServiceRequestDTO describes a service request payload.
type ServiceRequestDTO struct {
	RequestID      uint           `json:"request_id"`
	RequestDate    time.Time      `json:"request_date"`
	Description    string         `json:"description"`
	Status         string         `json:"status"`
	Priority       string         `json:"priority"`
	CompletionDate *time.Time     `json:"completion_date,omitempty"`
	TechnicianID   *uint          `json:"technician_id,omitempty"`
	Technician     *TechnicianDTO `json:"technician,omitempty"`
}
