package model

type EmployeeDTO struct {
	EmployeeID uint     `json:"employee_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	JobTitle   JobTitle `json:"job_title"`
	Gender     Gender   `json:"gender"`
	Email      string   `json:"email"`
	Salary     float64  `json:"salary"`
	ManagerID  *uint    `json:"manager_id,omitempty"`
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
	SaleID               uint        `json:"sale_id"`
	SaleDate             string      `json:"sale_date"`
	Amount               float64     `json:"amount"`
	SaleRepresentativeID uint        `json:"sale_representative_id"`
	SaleRepresentative   EmployeeDTO `json:"sale_representative,omitempty"`
	CustomerID           uint        `json:"customer_id"`
	Customer             CustomerDTO `json:"customer,omitempty"`
}

// ServiceRequestDTO describes a service request payload.
type ServiceRequestDTO struct {
	RequestID    uint        `json:"request_id"`
	Description  string      `json:"description"`
	TechnicianID uint        `json:"technician_id"`
	Technician   EmployeeDTO `json:"technician,omitempty"`
	CustomerID   uint        `json:"customer_id"`
	Customer     CustomerDTO `json:"customer,omitempty"`
}
