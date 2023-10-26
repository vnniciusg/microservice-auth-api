package userrepository

import (
	"database/sql"
	"errors"

	"github.com/vnniciusg/microservice-auth-api/internal/dto"
)

func (userRepository UserRepository) Delete(userRequest *dto.DeleteUserRequest) error {

	query := "SELECT * FROM user WHERE id = $1"
	err := userRepository.db.QueryRow(query, userRequest.ID).Scan(new(int))

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	deleteQuery := "DELETE FROM user WHERE id = $1"
	_, err = userRepository.db.Exec(deleteQuery, userRequest.ID)

	if err != nil {
		return err
	}

	return nil
}
