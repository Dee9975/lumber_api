package data

import "time"

type WorkdayRaw struct {
	ID          int       `json:"id"`
	TeamID      int       `json:"team_id"`
	WarehouseID int       `json:"warehouse_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Workday struct {
	Day       time.Time            `json:"day"`
	Team      Team                 `json:"team"`
	Warehouse Warehouse            `json:"warehouse"`
	Lumber    []TeamLumberResponse `json:"lumber"`
}

type Worker struct {
	Employee            Employee `json:"employee"`
	HoursWorked         int      `json:"hours_worked"`
	HoursWorkedForklift int      `json:"hours_worked_forklift"`
	HoursWorkedHeating  int      `json:"hours_worked_heating"`
}

type Team struct {
	ID        int                    `json:"id"`
	Employees []TeamEmployeeResponse `json:"employees"`
}

type Warehouse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Rate       int    `json:"rate"`
	CustomRate int    `json:"custom_rate"`
}
