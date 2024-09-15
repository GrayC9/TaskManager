package models

type TaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SeverityID  int    `json:"severity_id"`
	EmployeeID  int    `json:"employee_id"`
}
