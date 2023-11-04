package main

import (
	"context"
	"os"

	"github.com/martinyonatann/auth-microservices-traefik/auth-service/config"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/app"
)

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(context.Background(), cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
