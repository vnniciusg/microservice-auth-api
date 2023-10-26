package userrepository

import (
	"github.com/vnniciusg/microservice-auth-api/internal/domain"
)

func (userRepository UserRepository) Fetch() ([]*domain.User, error) {

	users := []*domain.User{}

	query := "SELECT * FROM user"

	rows, err := userRepository.db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &domain.User{}

		rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)

		users = append(users, user)
	}

	return users, nil
}
