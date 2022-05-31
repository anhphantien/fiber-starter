package handlers

import (
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/env"
	"fiber-starter/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type Claims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// @Summary Sign in
// @Tags auth
// @Param body body dto.SignInBody true " "
// @Success 200 {object} HttpResponse{data=SignInResponse}
// @Router /v1/auth/signin [post]
func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	body := dto.SignInBody{}
	if err, ok := Validate(c, &body); !ok {
		return err
	}

	db := database.DBConn

	user := models.User{}
	if err := db.Model(&models.User{}).First(&user, models.User{Username: &body.Username}).Error; err != nil {
		return SqlError(c, err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(body.Password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return c.Status(fiber.StatusBadRequest).JSON(HttpResponse{
				StatusCode: fiber.StatusBadRequest,
				Error:      INVALID_PASSWORD,
			})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(HttpResponse{
				StatusCode: fiber.StatusBadRequest,
				Error:      err.Error(),
			})
		}
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:       user.ID,
		Username: *user.Username,
		Role:     *user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(env.JWT_EXPIRES_AT) * time.Second)),
		},
	}).SignedString(env.JWT_SECRET)

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       SignInResponse{AccessToken: token},
	})
}
