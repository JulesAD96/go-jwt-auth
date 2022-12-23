package services

import "goauth/jwt/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
