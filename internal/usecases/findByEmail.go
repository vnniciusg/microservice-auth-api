package usecases

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

func (userUseCase UserUseCase) FindByEmail(userRequest *dto.FindByEmailRequest) (*dto.FindByEmailResponse, error) {
	user, err := userUseCase.UserRepository.FindByEmail(userRequest)

	if err != nil {
		return nil, err
	}
	return user, nil
}
