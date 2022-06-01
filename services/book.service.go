package services

import (
	"fiber-starter/common"
	"fiber-starter/database"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type BookService struct{}

// @Summary Get all books
// @Tags books
// @Success 200 {object} common.HttpResponse{data=[]entities.Book}
// @Router /v1/books [get]
func (h BookService) GetAll(c *fiber.Ctx) error {
	db := database.DB

	var books = entities.Book{}

	if err := db.Model(&entities.Book{}).Find(&books).Error; err != nil {
		return errors.SqlError(c, err)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       books,
	})
}

// @Summary Get a book by ID
// @Tags books
// @Param id path int true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books/{id} [get]
func (h BookService) GetByID(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	book := entities.Book{}

	if err := db.Model(&entities.Book{}).First(&book, id).Error; err != nil {
		return errors.SqlError(c, err)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       book,
	})
}

// @Summary Create a new book
// @Tags books
// @Param body body entities.Book true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books [post]
func (h BookService) Create(c *fiber.Ctx) error {
	db := database.DB

	book := entities.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	db.Create(&book)

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Summary Update a book
// @Tags books
// @Param id path int true " "
// @Param body body entities.Book true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books/{id} [put]
func (h BookService) Update(c *fiber.Ctx) error {
	db := database.DB

	book := entities.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	db.Create(&book)

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Security BearerAuth
// @Summary Delete a book
// @Tags books
// @Param id path int true " "
// @Success 200 {object} common.HttpResponse{}
// @Router /v1/books/{id} [delete]
func (h BookService) Delete(c *fiber.Ctx) error {
	user := utils.CurrentUser(c)
	if err, ok := utils.RoleAuth(c, user, []string{"ADMIN", "USER"}); !ok {
		return err
	}

	db := database.DB

	id := c.Params("id")
	book := entities.Book{}

	if err := db.First(&book, id).Error; err != nil {
		return errors.SqlError(c, err)
	}

	db.Delete(&book)

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
	})
}
