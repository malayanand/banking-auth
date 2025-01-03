package service

import (
	"github.com/malayanand/banking-auth/domain"
	"github.com/malayanand/banking-auth/dto"
	"github.com/malayanand/banking/errs"
)

type DefaultAuthService struct {
	repo domain.AuthRepository
}

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	var login *domain.Login
	var appErr *errs.AppError

	// get the user details. If the user exists
	if login, appErr = s.repo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	// get login claims from domain
	claims := login.ClaimsForAccessToken()
	// get auth token using the login claims just fetched
	authToken := domain.NewAuthToken(claims)

	var accessToken, refreshToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	if refreshToken, appErr = s.repo.GenerateAndSaveRefreshTokenToStore(authToken); appErr != nil {
		return nil, appErr
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
