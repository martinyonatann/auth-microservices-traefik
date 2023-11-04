package app

import (
	userV1 "github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/delivery/http/v1"
	userRepository "github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/repository"
	userUseCase "github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/usecase"
)

func (app *App) StartService() error {
	// define repository
	userRepo := userRepository.NewRepository(app.DB, app.Log)

	// define usecase
	userUC := userUseCase.NewUseCase(userRepo, app.Log, app.Cfg)

	// define controllers
	userCTRL := userV1.NewHandlers(userUC, app.Log)

	version := app.Echo.Group("/api/v1/")

	userV1.UserPrivateRoute(version, userCTRL, app.Cfg)

	return nil
}
