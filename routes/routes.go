package routes

import (
	"miniProject/config"
	"miniProject/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	NewUserControllers(e)
	NewArrivesControllers(e)
	NewDeliveriesControllers(e)
	NewProcessControllers(e)

	return e
}

func NewUserControllers(e *echo.Echo) {
	secured := e.Group("")
	secured.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	//Auth
	secured.GET("users", controllers.GetUsersController)
	secured.GET("/users/:id", controllers.GetUserController)
	secured.DELETE("/users/:id", controllers.DeleteUserController)
	secured.PUT("/users/:id", controllers.UpdateUserController)
}

func NewArrivesControllers(e *echo.Echo) {
	e.GET("arrives", controllers.GetArrivesController)
	e.GET("arrives/:id", controllers.GetArrivesIDController)
	e.POST("arrives", controllers.CreateArrivesController)
	e.DELETE("arrives/:id", controllers.DeleteArrivesController)
	e.PUT("arrives/:id", controllers.UpdateArrivesController)
}

func NewProcessControllers(e *echo.Echo) {
	e.GET("process", controllers.GetProcessController)
	e.GET("process/:id", controllers.GetProcessIDController)
	e.POST("process", controllers.CreateProcessController)
	e.DELETE("process/:id", controllers.DeleteProcessController)
	e.PUT("process/:id", controllers.UpdateProcessController)
}

func NewDeliveriesControllers(e *echo.Echo) {
	e.GET("deliveries", controllers.GetDeliveriesController)
	e.GET("deliveries/:id", controllers.GetDeliveriesIDController)
	e.POST("deliveries", controllers.CreateDeliveriesController)
	e.DELETE("deliveries/:id", controllers.DeleteDeliveriesController)
	e.PUT("deliveries/:id", controllers.UpdateDeliveriesController)
}
