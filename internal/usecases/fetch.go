package usecases

import (
	"github.com/vnniciusg/microservice-auth-api/internal/domain"
)

func (userUseCase UserUseCase) Fetch() ([]*domain.User, error) {
	users, err := userUseCase.UserRepository.Fetch()

	if err != nil {
		return nil, err
	}

	return users, nil
}
