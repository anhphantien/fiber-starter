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

// @Summary Sign in
// @Tags auth
// @Param body body dto.LoginBody true " "
// @Success 200 {object} common.HttpResponse{data=models.LoginResponse}
// @Router /v1/auth/login [post]
func (h AuthService) Login(c *fiber.Ctx) error {
	body := dto.LoginBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	db := database.DB

	user := entities.User{}
	if err := db.
		Model(&entities.User{}).
		First(&user, entities.User{Username: &body.Username}).Error; err != nil {
		return errors.SqlError(c, err)
	}
	if err := bcrypt.
		CompareHashAndPassword([]byte(*user.HashedPassword), []byte(body.Password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
				StatusCode: fiber.StatusBadRequest,
				Error:      errors.INVALID_PASSWORD,
			})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
				StatusCode: fiber.StatusBadRequest,
				Error:      err.Error(),
			})
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
