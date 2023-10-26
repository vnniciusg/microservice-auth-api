package userrepository

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

func (userRepository UserRepository) Create(userRequest *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {

	user := &dto.CreateUserResponse{}

	query := `
		INSERT INTO user (id, firstname , lastname, email, password) VALUES ($1 , $2 , $3 , $4 , $5)
	`

	err := userRepository.db.QueryRow(query, userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
