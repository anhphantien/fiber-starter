package handlers

import (
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// @Summary Sign in
// @Tags auth
// @Param Body body dto.SignInBody true " "
// @Success 200 {object} HttpResponse{data=SignInResponse}
// @Router /v1/auth/signin [post]
func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	signInBody := new(dto.SignInBody)
	if err, ok := Validate(c, signInBody); !ok {
		return err
	}

	db := database.DBConn

	var books = []_Book{}

	if err := db.Model(&models.Book{}).Find(&books).Error; err != nil {
		return SqlError(c, err)
	}

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       &books,
	})
}
