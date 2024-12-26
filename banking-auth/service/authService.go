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
	//var login *domain.Login

	if _, appErr := s.repo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	return nil, nil
}
