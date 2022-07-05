package repositories

import (
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

var user = entities.User{}

type UserRepository struct{}

func (r UserRepository) FindOneByID(c *fiber.Ctx, id any) (entities.User, error, bool) {
	err := CreateSqlBuilder(user).
		Where("id = ?", utils.ConvertToID(id)).
		Take(&user).Error
	if err != nil {
		return user, errors.SqlError(c, err), false
	}
	return user, nil, true
}
