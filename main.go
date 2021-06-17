package main

import (
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tiago123456789/api-golang/repositories"
)


func main() {
	e := echo.New()

	userRepository := repositories.NewUserRepository()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.POST("/users", userRepository.CreateUser)
	e.GET("/users/:id", userRepository.GetUserById)
	e.PUT("/users/:id", userRepository.UpdateUser)
	e.GET("/users", userRepository.GetUsers)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
