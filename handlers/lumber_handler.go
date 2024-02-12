package handlers

import (
	"github.com/labstack/echo/v4"
	"lumber/data"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetLumber(ctx echo.Context) error {
	lumber, err := h.lumberStore.GetLumber()

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, lumber)

	return nil
}

func (h *Handler) HandleCreateLumber(ctx echo.Context) error {
	var l data.Lumber

	err := ctx.Bind(&l)
	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing lumber data"})
		return err
	}

	if err := h.lumberStore.CreateLumber(l); err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusCreated, echo.Map{"success": true})
	return nil
}

func (h *Handler) HandleGetLumberById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	response, err := h.lumberStore.GetLumberById(id)

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, response)

	return nil
}

func (h *Handler) HandleDeleteLumber(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	if err := h.lumberStore.DeleteLumber(id); err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, echo.Map{"success": true})
	return nil
}

func (h *Handler) HandleUpdateLumber(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	var l data.Lumber

	err = ctx.Bind(&l)

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing lumber data"})
		return err
	}

	response, err := h.lumberStore.UpdateLumber(id, l)

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}

	_ = ctx.JSON(http.StatusOK, response)
	return nil
}

func (h *Handler) HandleGetAllTeamLumber(ctx echo.Context) error {
	teamID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		_ = ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id parameter"})
		return err
	}

	lumber, err := h.lumberStore.GetAllTeamLumber(teamID)

	if err != nil {
		_ = ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return err
	}

	_ = ctx.JSON(http.StatusOK, lumber)

	return nil
}
