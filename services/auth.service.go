package services

import "goauth/jwt/models"

type AuthService interface {
	SignUpUser(*models.SignInInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
