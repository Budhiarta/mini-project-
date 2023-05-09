package controllers

import (
	"net/http"
	"strconv"

	"miniProject/database"
	"miniProject/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetArrivesController(c echo.Context) error {
	arrives, err := database.GetArrives(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Arrive",
		"data":    arrives,
	})
}

func GetArrivesIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arrive, err := database.GetArrivesByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting arrive",
		"data":    arrive,
	})
}

func CreateArrivesController(c echo.Context) error {
	arrive := models.Arrive{}
	c.Bind(&arrive)

	arrive.Tot_price = arrive.Count * arrive.Price

	newArrive, err := database.CreateArrives(c.Request().Context(), arrive)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating arrive",
		"data":    newArrive,
	})
}

func DeleteArrivesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteArrives(c.Request().Context(), id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting arrive data",
	})
}

func UpdateArrivesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arrives := models.Arrive{}
	c.Bind(&arrives)

	updatedArrive, err := database.UpdateArrives(c.Request().Context(), id, arrives)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating arrive data",
		"data":    updatedArrive,
	})
}
