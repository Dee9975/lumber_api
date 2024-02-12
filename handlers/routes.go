package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(v1 *echo.Group) {
	v1.GET("/lumber", h.HandleGetLumber)
	v1.POST("/lumber", h.HandleCreateLumber)
	v1.GET("/lumber/:id", h.HandleGetLumberById)
	v1.DELETE("/lumber/:id", h.HandleDeleteLumber)
	v1.PUT("/lumber/:id", h.HandleUpdateLumber)
	v1.GET("/lumber/team/:id", h.HandleGetAllTeamLumber)

	v1.GET("/employees", h.HandleGetAllEmployees)
	v1.GET("/employees/:id", h.HandleGetEmployeeById)
	v1.POST("/employees", h.HandleCreateEmployee)
	v1.PUT("/employees/:id", h.HandleUpdateEmployee)
	v1.DELETE("/employees/:id", h.HandleDeleteEmployee)
	v1.GET("/employees/team/:id", h.HandleGetAllEmployeesByTeamId)

	v1.GET("/workdays", h.HandleGetAllWorkdays)
}
