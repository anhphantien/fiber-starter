package handlers

import (
	"fiber-starter/database"
	"fiber-starter/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct{}

type _Book struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

// @Summary Get all books
// @Tags books
// @Success 200 {object} HttpResponse{data=[]models.Book}
// @Router /v1/books [get]
func (h BookHandler) GetAll(c *fiber.Ctx) error {
	db := database.DB

	var books = []_Book{}

	if err := db.Model(&models.Book{}).Find(&books).Error; err != nil {
		return SqlError(c, err)
	}

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       books,
	})
}

// @Summary Get a book by ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} HttpResponse{data=models.Book}
// @Router /v1/books/{id} [get]
func (h BookHandler) GetByID(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	book := _Book{}

	if err := db.Model(&models.Book{}).First(&book, id).Error; err != nil {
		return SqlError(c, err)
	}

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       book,
	})
}

// @Summary Create a new book
// @Tags books
// @Param body body models.Book true " "
// @Success 200 {object} HttpResponse{data=models.Book}
// @Router /v1/books [post]
func (h BookHandler) Create(c *fiber.Ctx) error {
	db := database.DB

	book := models.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	db.Create(&book)

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Security BearerAuth
// @Summary Delete a book
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} HttpResponse{}
// @Router /v1/books/{id} [delete]
func (h BookHandler) Delete(c *fiber.Ctx) error {
	user := CurrentUser(c)
	fmt.Println(user)

	db := database.DB

	id := c.Params("id")
	book := models.Book{}

	if err := db.First(&book, id).Error; err != nil {
		return SqlError(c, err)
	}

	db.Delete(&book)

	return c.JSON(HttpResponse{
		StatusCode: fiber.StatusOK,
	})
}
