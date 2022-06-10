package services

import (
	"fiber-starter/common"
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/env"
	"fiber-starter/errors"
	"fiber-starter/models"
	"fiber-starter/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

// @Summary Login
// @Tags auth
// @Param body body dto.LoginBody true " "
// @Success 200 {object} common.HttpResponse{data=models.LoginResponse}
// @Router /v1/auth/login [post]
func (s AuthService) Login(c *fiber.Ctx) error {
	db := database.DB

	body := dto.LoginBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	user := entities.User{}
	r := db.
		Model(&user).
		Where("username = ?", body.Username).
		First(&user)
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

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		ID:       user.ID,
		Username: *user.Username,
		Role:     *user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().
				Add(time.Duration(env.JWT_EXPIRES_AT) * time.Second)),
		},
	}).SignedString(env.JWT_SECRET)

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       models.LoginResponse{AccessToken: token},
	})
}
