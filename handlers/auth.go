package handlers

import (
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/models"
	"fiber-starter/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// @Summary Sign in
// @Tags auth
// @Param Body body dto.SignIn true " "
// @Success 200 {object} response.Http{data=models.User}
// @Router /v1/auth/signin [post]
func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	signIn := new(dto.SignIn)
	response.Validate(c, signIn)

	db := database.DBConn

	var books = []_Book{}

	if err := db.Model(&models.Book{}).Find(&books).Error; err != nil {
		return response.Error(c, err)
	}

	return c.JSON(response.Http{
		StatusCode: http.StatusOK,
		Data:       &books,
	})
}
