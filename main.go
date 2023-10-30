package main

import (
	"context"
	"fmt"
	"log"
	"m2ps/app"
	"m2ps/config"
	"m2ps/helpers"
	"m2ps/repositories"
	"m2ps/routes"
	"m2ps/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Passing Variable
var (
	echoHandler echo.Echo
	ctx         = context.Background()
)

func main() {
	if err := config.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open Connection Faild: %s", err.Error()))
	}
	defer config.CloseConnectionDB()

	// Mongo DB connection using database default
	mongoDB := config.ConnectMongo(ctx)
	defer config.CloseMongo(ctx)

	// Connection database
	DB := config.DBConnection()
	//DBCloud := config.DBConnectionCloud()
	log.Println("DB CONNECTION : ", utils.ToString(DB))

	// Configuration Repository
	repo := repositories.NewRepository(DB, mongoDB, ctx, nil)

	// Configuration Repository and Services

	echoHandler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	echoHandler.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		result := helpers.ResponseJSON(false, strconv.Itoa(report.Code), err.Error(), nil)
		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}

	services := app.SetupApp(DB, repo)

	// Routing API
	routes.RoutesApi(echoHandler, services)

	port := fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8080"))
	echoHandler.Logger.Fatal(echoHandler.Start(port))
}
