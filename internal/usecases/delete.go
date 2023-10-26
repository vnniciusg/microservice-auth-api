package usecases

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

func (userUseCase UserUseCase) Delete(userRequest *dto.DeleteUserRequest) error {
	err := userUseCase.UserRepository.Delete(userRequest)
	if err != nil {
		return err
	}

	return nil
}
