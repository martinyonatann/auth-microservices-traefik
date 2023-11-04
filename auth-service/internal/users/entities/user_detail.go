package entities

import "github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/dtos"

func NewUserDetail(data Users) dtos.UserDetailResponse {
	return dtos.UserDetailResponse{
		UserID:      data.UserID,
		Email:       data.Email,
		Fullname:    data.Fullname,
		PhoneNumber: data.PhoneNumber,
		UserType:    data.UserType,
		IsActive:    data.IsActive,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy,
	}
}
