package dtos

import (
	"github.com/invopop/validation"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/apperror"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/constant"
)

type UpdateUserRequest struct {
	UserID      int64  `json:"-"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UpdateBy    string `json:"-"`
}

var ()

func (cup UpdateUserRequest) Validate() error {
	if cup.UserType != "" && cup.UserType != constant.UserTypePremium && cup.UserType != constant.UserTypeRegular {
		return apperror.ErrInvalidUserType
	}

	return validation.ValidateStruct(&cup,
		validation.Field(&cup.UserID, validation.NotNil),
		validation.Field(&cup.UpdateBy, validation.NotNil),
	)
}
