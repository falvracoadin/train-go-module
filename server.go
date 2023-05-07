package main

import (
	"go-skeleton-manager-rabbitmq/helpers"
	"go-skeleton-manager-rabbitmq/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	errLoadingEnvFile := godotenv.Load()
	if errLoadingEnvFile != nil {
		helpers.HandleError("error loading the .env file", errLoadingEnvFile)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	routes.Build(e)
	//routes.RouteNonAuth(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
