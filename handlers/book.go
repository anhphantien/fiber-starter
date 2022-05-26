package handlers

import (
	"fiber-starter/database"
	"fiber-starter/models"
	"fiber-starter/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct{}

type _Book struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

// @Summary Get all books
// @Tags books
// @Success 200 {object} response.Http{data=[]models.Book}
// @Router /v1/books [get]
func (h BookHandler) GetAll(c *fiber.Ctx) error {
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

// @Summary Get a book by ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} response.Http{data=models.Book}
// @Router /v1/books/{id} [get]
func (h BookHandler) GetByID(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	book := _Book{}

	if err := db.Model(&models.Book{}).First(&book, id).Error; err != nil {
		return response.Error(c, err)
	}

	return c.JSON(response.Http{
		StatusCode: http.StatusOK,
		Data:       &book,
	})
}

// @Summary Create a new book
// @Tags books
// @Param Body body models.Book true " "
// @Success 200 {object} response.Http{data=models.Book}
// @Router /v1/books [post]
func (h BookHandler) Create(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Http{
			StatusCode: http.StatusBadRequest,
		})
	}

	db.Create(book)

	return c.JSON(response.Http{
		StatusCode: http.StatusCreated,
		Data:       *book,
	})
}

// @Security BearerAuth
// @Summary Delete a book
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} response.Http{}
// @Router /v1/books/{id} [delete]
func (h BookHandler) Delete(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	book := new(models.Book)

	if err := db.First(&book, id).Error; err != nil {
		return response.Error(c, err)
	}

	db.Delete(&book)

	return c.JSON(response.Http{
		StatusCode: http.StatusOK,
	})
}
