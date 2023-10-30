package routes

import (
	"m2ps/config"
	"m2ps/services"
	"m2ps/services/trxService"

	"github.com/labstack/echo"
)

// RoutesApi
func RoutesApi(e echo.Echo, usecaseSvc services.UsecaseService) {
	routePrefix := e.Group(config.GetEnv("PREFIX_API"))

	trxSvc := trxService.NewTrxService(usecaseSvc)
	trxGroup := routePrefix.Group("/trx")
	trxGroup.POST("/add-trx", trxSvc.Create)

}
