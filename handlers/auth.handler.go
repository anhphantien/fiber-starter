package handlers

import (
	"fiber-starter/common"
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/env"
	"fiber-starter/errors"
	"fiber-starter/models"
	"fiber-starter/repositories"
	"fiber-starter/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

// @Tags    auth
// @Summary Login
// @Param   body           body       dto.LoginBody true " "
// @Success 201            {object}   common.Response{data=models.LoginResponse}
// @Router  /v1/auth/login [post]
func (h AuthHandler) Login(c *fiber.Ctx) error {
	body := dto.LoginBody{}
	if err, ok := utils.ValidateRequestBody(c, &body); !ok {
		return err
	}

	user := entities.User{}
	r := repositories.
		CreateSqlBuilder(user).
		Where("username = ?", body.Username).
		Take(&user)
	if r.Error != nil {
		return errors.SqlError(c, r.Error)
	}
	if err := bcrypt.
		CompareHashAndPassword(
			[]byte(*user.HashedPassword),
			[]byte(body.Password),
		); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return errors.BadRequestException(c, errors.INVALID_PASSWORD)
		default:
			return errors.BadRequestException(c, err.Error())
		}
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.CurrentUser{
			ID:       user.ID,
			Username: *user.Username,
			Role:     *user.Role,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().
				Add(time.Duration(env.JWT_EXPIRES_AT) * time.Second)),
		},
	).SignedString(env.JWT_SECRET)

	return common.HttpResponse(c, (common.Response{
		Data: models.LoginResponse{
			AccessToken: token,
		},
	}))
}
