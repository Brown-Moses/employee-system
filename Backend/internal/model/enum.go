package model

type JobTitle string

const (
	JobManager    JobTitle = "Manager"
	JobSales      JobTitle = "Sales Representative"
	JobTechnician JobTitle = "Technician"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)
