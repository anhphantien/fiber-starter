package handlers

import (
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/env"
	"fiber-starter/errors"
	"fiber-starter/models"
	"fiber-starter/repositories"
	"fiber-starter/response"
	"fiber-starter/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

// @Tags    auth
// @Summary Login
// @Param   body               body   dto.LoginBody true " "
// @Success 201                object response.Response{data=models.LoginResponse}
// @Router  /api/v1/auth/login [POST]
func (h AuthHandler) Login(c *fiber.Ctx) error {
	body := dto.LoginBody{}
	if err, ok := utils.ValidateRequestBody(c, &body); !ok {
		return err
	}

	user := entities.User{}
	err := repositories.
		CreateSqlBuilder(user).
		Where("username = ?", body.Username).
		Take(&user).Error
	if err != nil {
		return errors.SqlError(c, err)
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
			ID:        user.ID,
			Username:  *user.Username,
			Role:      *user.Role,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(env.JWT_EXPIRES_AT * time.Second).Unix(),
		},
	).SignedString(env.JWT_SECRET)

	return response.WriteJSON(c, (response.Response{
		Data: models.LoginResponse{
			AccessToken: token,
		},
	}))
}
