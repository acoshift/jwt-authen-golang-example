package main

import (
	"net/http"
	"time"

	"jwt-authen-golang-example/api"
	"jwt-authen-golang-example/service"
	"log"

	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

const projectID = "jwt-authen-example"

func main() {
	serviceAccount, err := ioutil.ReadFile("service-account.json")
	if err != nil {
		log.Fatal(err)
	}
	err = api.Init(api.Config{
		ServiceAccountJSON: serviceAccount,
		ProjectID:          projectID,
	})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.BodyLimit("2M"),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:8080",
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderContentType,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				echo.GET,
				echo.POST,
			},
			MaxAge: 3600,
		}),
	)

	// Health check
	e.Get("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Register services
	service.Auth(e.Group("/auth"))

	e.Run(standard.WithConfig(engine.Config{
		Address:      ":9000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}))
}
