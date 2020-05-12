package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zvezduk/go/pkg/api"
	"github.com/zvezduk/go/pkg/resource"
	"gopkg.in/go-playground/validator.v9"
	"os"
)

type config struct {
	port    string
	storage string
}

var conf config

func init() {
	godotenv.Overload()

	conf = config{
		port:    os.Getenv("PORT"),
		storage: os.Getenv("STORAGE_URL"),
	}

	resource.StorageUrl = conf.storage
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.Validator = &api.RequestValidator{Validator: validator.New()}

	api.Public(e)

	e.Logger.Fatal(e.Start(conf.port))
}
