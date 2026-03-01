package model

type JobTitle string

const (
	JobEngineer JobTitle = "Engineer"
	JobManager  JobTitle = "Manager"
	JobSales    JobTitle = "Sales Representative"
	JobTech     JobTitle = "Technician"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)
