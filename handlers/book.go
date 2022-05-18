package handlers

import (
	"fiber-starter/database"
	"fiber-starter/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}

// @Summary Get all books
// @Tags books
// @Success 200 {object} HttpResponse{data=[]models.Book}
// @Router /v1/books [get]
func GetAll(c *fiber.Ctx) error {
	db := database.DBConn

	var books []models.Book
	if res := db.Find(&books); res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(HttpResponse{
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(HttpResponse{
		StatusCode: http.StatusOK,
		Data:       books,
	})
}

// @Summary Get a book by ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} HttpResponse{data=models.Book}
// @Router /v1/books/{id} [get]
func GetByID(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	book := new(models.Book)
	if err := db.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(http.StatusNotFound).JSON(HttpResponse{
				StatusCode: http.StatusNotFound,
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(HttpResponse{
				StatusCode: http.StatusInternalServerError,
			})
		}
	}

	return c.JSON(HttpResponse{
		StatusCode: http.StatusOK,
		Data:       *book,
	})
}

// @Summary Create a new book
// @Tags books
// @Param book body models.Book true "Book data"
// @Success 200 {object} HttpResponse{data=models.Book}
// @Router /v1/books [post]
func Create(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(HttpResponse{
			StatusCode: http.StatusBadRequest,
		})
	}

	db.Create(book)

	return c.JSON(HttpResponse{
		StatusCode: http.StatusCreated,
		Data:       *book,
	})
}

// @Security BearerAuth
// @Summary Delete a book
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} HttpResponse{}
// @Router /v1/books/{id} [delete]
func Delete(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	book := new(models.Book)
	if err := db.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(http.StatusNotFound).JSON(HttpResponse{
				StatusCode: http.StatusNotFound,
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(HttpResponse{
				StatusCode: http.StatusInternalServerError,
			})
		}
	}

	db.Delete(&book)

	return c.JSON(HttpResponse{
		StatusCode: http.StatusNoContent,
	})
}
