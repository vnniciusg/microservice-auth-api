package handler

import "github.com/vnniciusg/microservice-auth-api/internal/domain"

type UserService struct {
	UserUseCase domain.UserUseCase
}

func New(userUseCase domain.UserUseCase) *UserService {
	return &UserService{
		UserUseCase: userUseCase,
	}
}
