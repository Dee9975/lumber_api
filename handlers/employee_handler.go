package handlers

import (
	"github.com/labstack/echo/v4"
	"lumber/data"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetAllEmployees(ctx echo.Context) error {
	employees, err := h.employeeStore.GetAllEmployees()

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, employees)

	return nil
}

func (h *Handler) HandleGetEmployeeById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	employee, err := h.employeeStore.GetEmployeeById(id)

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, employee)

	return nil
}

func (h *Handler) HandleCreateEmployee(ctx echo.Context) error {
	var e data.Employee

	err := ctx.Bind(&e)
	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing employee data"})
		return err
	}

	if err := h.employeeStore.CreateEmployee(e); err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusCreated, echo.Map{"success": true})
	return nil
}

func (h *Handler) HandleUpdateEmployee(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	var e data.Employee

	err = ctx.Bind(&e)

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing employee data"})
		return err
	}

	if err := h.employeeStore.UpdateEmployee(id, e); err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, echo.Map{"success": true})
	return nil
}

func (h *Handler) HandleDeleteEmployee(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	if err := h.employeeStore.DeleteEmployee(id); err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, echo.Map{"success": true})
	return nil
}

func (h *Handler) HandleGetAllEmployeesByTeamId(ctx echo.Context) error {
	teamID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing team_id parameter"})
		return err
	}

	employees, err := h.employeeStore.GetEmployeesByTeamId(teamID)

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, employees)

	return nil
}
