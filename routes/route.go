package routes

import (
	"telkomsel/controllers"
	"telkomsel/utils"

	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.ReadAllUsers)
	e.GET("/users/:id", controllers.GetUserByID)
	e.PUT("/users/:id", controllers.UpdateUserByID)
	e.DELETE("/users/:id", controllers.DeleteUserByID)
	e.GET("/pictures/:url", utils.CheckPicture)

	e.Static("/", "/views")
	e.Static("/assets", "/assets")

	return e
}
