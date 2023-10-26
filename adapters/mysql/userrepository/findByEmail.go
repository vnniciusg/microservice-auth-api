package userrepository

import (
	"github.com/vnniciusg/microservice-auth-api/internal/domain"
	"github.com/vnniciusg/microservice-auth-api/internal/dto"
)

func (userRepository UserRepository) FindByEmail(userRequest *dto.FindByEmailRequest) (*domain.User, error) {

	user := &domain.User{}

	query := `SELECT * FROM user WHERE email = $1`

	err := userRepository.db.QueryRow(query, userRequest.Email).Scan(&user.ID, &user.FirstName, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
