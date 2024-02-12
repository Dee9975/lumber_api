package handlers

import (
	"github.com/labstack/echo/v4"
	"lumber/data"
	"net/http"
)

func (h *Handler) HandleGetAllWorkdays(ctx echo.Context) error {
	date := ctx.QueryParam("day")

	if date != "" {
		rawWorkdays, err := h.workdayStore.GetWorkdaysFromDate(date)
		if err != nil {
			_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning raw workdays: " + err.Error()})
			return err
		}
		workdays := []data.Workday{}
		for _, rawWorkday := range rawWorkdays {
			employees, err := h.employeeStore.GetEmployeesByTeamId(rawWorkday.TeamID)
			if err != nil {
				_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning employees: " + err.Error()})
				return err
			}
			warehouse, err := h.lumberStore.GetWarehouseById(rawWorkday.WarehouseID)
			if err != nil {
				_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning warehouse: " + err.Error()})
				return err
			}
			lumber, err := h.lumberStore.GetAllTeamLumber(rawWorkday.TeamID)
			if err != nil {
				_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning lumber: " + err.Error()})
				return err
			}
			workday := data.Workday{
				Day:       rawWorkday.CreatedAt,
				Team:      data.Team{ID: rawWorkday.TeamID, Employees: employees},
				Lumber:    lumber,
				Warehouse: *warehouse,
			}
			workdays = append(workdays, workday)
		}
		_ = ctx.JSON(http.StatusOK, workdays)
		return nil
	}

	rawWorkdays, err := h.workdayStore.GetAllWorkdays()

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning raw workdays: " + err.Error()})
		return err
	}

	workdays := []data.Workday{}
	for _, rawWorkday := range rawWorkdays {
		employees, err := h.employeeStore.GetEmployeesByTeamId(rawWorkday.TeamID)
		if err != nil {
			_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning employees: " + err.Error()})
			return err
		}

		warehouse, err := h.lumberStore.GetWarehouseById(rawWorkday.WarehouseID)
		if err != nil {
			_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning warehouse: " + err.Error()})
			return err
		}

		lumber, err := h.lumberStore.GetAllTeamLumber(rawWorkday.TeamID)

		if err != nil {
			_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning lumber: " + err.Error()})
			return err
		}

		workday := data.Workday{
			Day:       rawWorkday.CreatedAt,
			Team:      data.Team{ID: rawWorkday.TeamID, Employees: employees},
			Lumber:    lumber,
			Warehouse: *warehouse,
		}

		workdays = append(workdays, workday)
	}

	_ = ctx.JSON(http.StatusOK, workdays)
	return nil
}
