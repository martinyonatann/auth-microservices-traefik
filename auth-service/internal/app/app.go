package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/config"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/datasource"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/middleware"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/otel/zerolog"
)

type App struct {
	DB   *sqlx.DB
	Echo *echo.Echo
	Log  *zerolog.Logger
	Cfg  config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	return &App{
		DB:   db,
		Echo: echo.New(),
		Log:  zerolog.NewZeroLog(ctx, os.Stdout),
		Cfg:  cfg,
	}
}

func (app *App) Start() error {
	if err := app.StartService(); err != nil {
		app.Log.Z().Err(err).Msg("[app]StartService")

		return err
	}

	app.Echo.Debug = app.Cfg.Server.Debug
	app.Echo.Use(middleware.AppCors())

	return app.Echo.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%s", app.Cfg.Server.RESTPort),
		ReadTimeout:  app.Cfg.Server.ReadTimeout,
		WriteTimeout: app.Cfg.Server.WriteTimeout,
	})
}
