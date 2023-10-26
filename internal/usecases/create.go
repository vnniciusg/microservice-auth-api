package usecases

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

func (userUseCase UserUseCase) Create(userRequest *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	user, err := userUseCase.UserRepository.Create(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}
