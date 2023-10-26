package usecases

import "github.com/vnniciusg/microservice-auth-api/internal/domain"

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func New(userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}
