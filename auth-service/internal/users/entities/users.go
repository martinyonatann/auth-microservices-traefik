package entities

import (
	"time"

	"github.com/martinyonatann/auth-microservices-traefik/auth-service/config"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/dtos"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/constant"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/utils"
)

type (
	Users struct {
		UserID      int64
		Email       string
		Password    string
		Fullname    string
		PhoneNumber string
		UserType    string
		IsActive    bool
		CreatedAt   time.Time
		CreatedBy   string
		UpdatedAt   time.Time
		UpdatedBy   string
	}

	LockingOpt struct {
		PessimisticLocking bool
	}
)

func NewCreateUser(data dtos.CreateUserRequest, cfg config.Config) Users {
	return Users{
		Fullname:    data.FullName,
		Email:       data.Email,
		Password:    utils.Encrypt(data.Password, cfg),
		PhoneNumber: data.PhoneNumber,
		UserType:    constant.UserTypeRegular,
		IsActive:    true,
		CreatedAt:   time.Now(),
		CreatedBy:   "SYSTEM",
	}
}
