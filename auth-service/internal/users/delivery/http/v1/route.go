package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/config"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/middleware"
)

func UserPrivateRoute(version *echo.Group, h Handlers, cfg config.Config) {
	users := version.Group("users")
	users.POST("", h.CreateUser)
	users.POST("/login", h.Login)
	users.GET("/detail", h.UserDetail, middleware.AuthorizeJWT(cfg))
	users.PATCH("/update", h.UpdateUser, middleware.AuthorizeJWT(cfg))
	users.PUT("/update/status", h.UpdateUserStatus, middleware.AuthorizeJWT(cfg))
}
