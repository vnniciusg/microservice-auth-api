package usecases

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

func (userUseCase UserUseCase) Fetch() ([]*dto.FetchUser, error) {
	users, err := userUseCase.UserRepository.Fetch()

	if err != nil {
		return nil, err
	}

	return users, nil
}
