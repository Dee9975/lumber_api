package handlers

import "lumber/stores"

type Handler struct {
	lumberStore   *stores.LumberStore
	employeeStore *stores.EmployeeStore
	workdayStore  *stores.WorkdayStore
}

func NewHandler(lumberStore *stores.LumberStore, employeeStore *stores.EmployeeStore, workdayStore *stores.WorkdayStore) *Handler {
	return &Handler{
		lumberStore:   lumberStore,
		employeeStore: employeeStore,
		workdayStore:  workdayStore,
	}
}
