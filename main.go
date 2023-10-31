package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	id_translations "gopkg.in/go-playground/validator.v9/translations/id"
	"log"
	"m2ps/app"
	"m2ps/config"
	"m2ps/helpers"
	"m2ps/repositories"
	"m2ps/routes"
	"net/http"
)

//CustomValidator adalah
type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

// Passing Variable
var (
	uni         *ut.UniversalTranslator
	echoHandler echo.Echo
	ctx         = context.Background()
)

// Custom Validator and translation
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, row := range errs {
			return errors.New(row.Translate(cv.translator))
		}
	}

	return cv.validator.Struct(i)
}

func main() {
	if err := config.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open Connection Faild: %s", err.Error()))
	}
	defer config.CloseConnectionDB()

	//if err := config.OpenConnectionCloud(); err != nil {
	//	panic(fmt.Sprintf("Open Connection Faild: %s", err.Error()))
	//}
	//defer config.CloseConnectionDBCloud()

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
		result := helpers.ResponseJSON(false, "401", err.Error(), nil)
		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}

	// Mongo DB connection using database default
	mongoDB := config.ConnectMongo(ctx)
	defer config.CloseMongo(ctx)

	// Connection database
	DB := config.DBConnection()
	//DBCloud := config.DBConnectionCloud()

	//repoGenAutoNum := genautonum.NewRepository(nil, ctx, mongoDB)

	//Redis Connection
	//redis := config.ConnectRedis()
	//redisLocal := config.ConnectRedisLocal()

	// Configuration Repository
	repo := repositories.NewRepository(DB, mongoDB, ctx, nil)

	// Configuration Repository and Services
	services := app.SetupApp(DB, repo)

	echoHandler.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start] Request:", c.Path())
		log.Println("Request Body:", string(reqBody))
		log.Println("Response Body:", string(resBody))
		log.Println("[End] Request:", c.Path())
	}))

	// Routing API
	routes.RoutesApi(echoHandler, services)

	port := fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8080"))
	echoHandler.Logger.Fatal(echoHandler.Start(port))
}

func init() {
	boardingService()

	e := echo.New()
	echoHandler = *e
	validateCustom := validator.New()

	id := id.New()
	uni = ut.New(id, id)
	trans, _ := uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(validateCustom, trans)
	e.Validator = &CustomValidator{validator: validateCustom, translator: trans}

	e.Static("/img/*", "assets/img")
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	echoHandler.Use(middleware.Logger())
}

func boardingService() {
	fmt.Println(`
  __  __   _  __  _____      __  __ c          _      _   _             _____                         _               
 |  \/  | | |/ / |  __ \    |  \/  |         | |     (_) | |           / ____|                       (_)              
 | \  / | | ' /  | |__) |   | \  / |   ___   | |__    _  | |   ___    | (___     ___   _ __  __   __  _    ___    ___ 
 | |\/| | |  <   |  ___/    | |\/| |  / _ \  | '_ \  | | | |  / _ \    \___ \   / _ \ | '__| \ \ / / | |  / __|  / _ \
 | |  | | | . \  | |        | |  | | | (_) | | |_) | | | | | |  __/    ____) | |  __/ | |     \ V /  | | | (__  |  __/
 |_|  |_| |_|\_\ |_|        |_|  |_|  \___/  |_.__/  |_| |_|  \___|   |_____/   \___| |_|      \_/   |_|  \___|  \___|`)
}
