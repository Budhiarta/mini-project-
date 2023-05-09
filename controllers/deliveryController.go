package controllers

import (
	"net/http"
	"strconv"

	"miniProject/database"
	"miniProject/dto"
	"miniProject/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetDeliveriesController(c echo.Context) error {
	deliveries, err := database.GetDeliveries(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting deliveries data",
		"data":    deliveries,
	})
}

func GetDeliveriesIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	delivery, err := database.GetDeliveriesByID(c.Request().Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting deliveries data",
		"data":    delivery,
	})
}

func CreateDeliveriesController(c echo.Context) error {
	deliveries := dto.CreateDeliveryRequest{}

	c.Bind(&deliveries)

	for _, each := range deliveries.Proceses {
		err := database.DecreaseStock(c.Request().Context(), int(each.Id), each.Count)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	newdelivery, err := database.CreateDeliveries(c.Request().Context(), *deliveries.ToEntity())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating delivery",
		"data":    newdelivery,
	})
}

func DeleteDeliveriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteDeliveries(c.Request().Context(), id)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting delivery data",
	})
}

func UpdateDeliveriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	deliveries := models.Delivery{}
	c.Bind(&deliveries)

	updatedDelivery, err := database.UpdateDeliveries(c.Request().Context(), id, deliveries)
	if err != nil {
		if err == database.ErrInvalidID {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating delivery data",
		"data":    updatedDelivery,
	})
}
