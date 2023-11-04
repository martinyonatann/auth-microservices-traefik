package entities

import (
	"time"

	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/dtos"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/constant"
)

type UpdateUserStatus struct {
	UserID    int64
	IsActive  bool
	UpdatedAt time.Time
	UpdatedBy string
}

func NewUpdateUserStatus(req dtos.UpdateUserStatusRequest) UpdateUserStatus {
	return UpdateUserStatus{
		UserID:    req.UserID,
		IsActive:  constant.MapStatus[req.Status],
		UpdatedAt: time.Now(),
		UpdatedBy: req.UpdateBy,
	}
}
