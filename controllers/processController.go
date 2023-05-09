package controllers

import (
	"net/http"
	"strconv"

	"miniProject/database"
	"miniProject/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProcessController(c echo.Context) error {
	process, err := database.GetProcess(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Process data",
		"data":    process,
	})
}

func GetProcessIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	process, err := database.GetprocessByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting arrive",
		"data":    process,
	})
}

func CreateProcessController(c echo.Context) error {
	Process := models.Process{}
	c.Bind(&Process)

	NewProcess, err := database.Createprocess(c.Request().Context(), Process)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating process data",
		"data":    NewProcess,
	})

}

func DeleteProcessController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.Deleteprocess(c.Request().Context(), id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting process data",
	})
}

func UpdateProcessController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	process := models.Process{}
	c.Bind(&process)

	updatedProcess, err := database.Updateprocess(c.Request().Context(), id, process)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating arrive data",
		"data":    updatedProcess,
	})
}
