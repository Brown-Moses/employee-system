package model

type EmployeeDTO struct {
	EmployeeID uint     `json:"employee_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	JobTitle   JobTitle `json:"job_title"`
	Gender     Gender   `json:"gender"`
	Salary     float64  `json:"salary"`
	ManagerID  *uint    `json:"manager_id,omitempty"`
}
