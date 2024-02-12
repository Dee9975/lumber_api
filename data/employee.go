package data

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type EmployeeResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type TeamEmployeeResponse struct {
	ID                  int    `json:"id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Email               string `json:"email"`
	HoursWorked         int    `json:"hours_worked"`
	HoursWorkedForklift int    `json:"hours_worked_forklift"`
	HoursWorkedHeating  int    `json:"hours_worked_heating"`
}
