package repositories

import (
	"fiber-starter/database"
	"fiber-starter/entities"
)

type UserRepository struct{}

func (r UserRepository) FindByUsername(username string) (entities.User, error) {
	user := entities.User{}

	err := database.DB.
		Model(user).
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil {
		return user, err
	}

	return user, nil
}
