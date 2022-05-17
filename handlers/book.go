package handlers

import (
	"fiber-starter/database"
	"fiber-starter/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// @Summary Get all books
// @Tags books
// @Success 200 {object} ResponseHTTP{data=[]models.Book}
// @Router /v1/books [get]
func GetAllBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []models.Book
	if res := db.Find(&books); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all books.",
		Data:    books,
	})
}

// @Summary Get book by ID 11111111
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Book}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/books/{id} [get]
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Book with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get book by ID.",
		Data:    *book,
	})
}

// @Summary Register a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Register book"
// @Success 200 {object} ResponseHTTP{data=models.Book}
// @Failure 400 {object} ResponseHTTP{}
// @Router /v1/books [post]
func RegisterBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Create(book)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a book.",
		Data:    *book,
	})
}

// @Security BearerAuth
// @Summary Remove book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/books/{id} [delete]
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Book with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	db.Delete(&book)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success delete book.",
		Data:    nil,
	})
}
