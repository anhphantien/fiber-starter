package services

import (
	"database/sql"
	"fiber-starter/common"
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BookService struct{}

// @Tags    books
// @Summary Get a list of books
// @Param   limit     query    int false " "
// @Param   page      query    int false " "
// @Param   keyword   query    string false " "
// @Param   filter    query    object false " "
// @Param   sort      query    object false " "
// @Success 200       {object} common.HttpResponse{data=[]entities.Book}
// @Router  /v1/books [get]
func (s BookService) GetList(c *fiber.Ctx) error {
	db := database.DB

	books := []entities.Book{}

	pagination := utils.Pagination(c)

	q := db.
		Model(books).
		Joins("User") // LEFT JOIN (one-to-one)
		// Select("book.*" + utils.GetAllColumnsOfTable(entities.User{})).
		// Joins("INNER JOIN user ON book.user_id = user.id") // INNER JOIN (one-to-one)
	if pagination.Filter["id"] != nil {
		q.Where("id = ?", utils.ConvertToInt(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"title LIKE @keyword OR description LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}

	// var total int64
	// r1 := q.Count(&total)
	// if r1.Error != nil {
	// 	return errors.SqlError(c, r1.Error)
	// }

	// r2 := q.Limit(pagination.Limit).
	// 	Offset(pagination.Offset).
	// 	Order(pagination.Order).
	// 	Find(&books) // db.Table("book").Select("1 + 2 AS sum, \"abc\" AS title").Scan(&books)
	// if r2.Error != nil {
	// 	return errors.SqlError(c, r2.Error)
	// }

	ch := make(chan error, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var total int64
	go func() {
		defer wg.Done()

		r := q.
			Session(&gorm.Session{}). // clone
			Count(&total)
		if r.Error != nil {
			ch <- r.Error
		}
	}()

	go func() {
		defer wg.Done()

		r := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Order(pagination.Order).
			Find(&books)
		if r.Error != nil {
			ch <- r.Error
		}
	}()

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return errors.SqlError(c, err)
		}
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data: common.PaginationResponse{
			Items: books,
			Total: total,
		},
	})
}

// @Tags    books
// @Summary Get a book by ID
// @Param   id             path     int true " "
// @Success 200            {object} common.HttpResponse{data=entities.Book}
// @Router  /v1/books/{id} [get]
func (s BookService) GetByID(c *fiber.Ctx) error {
	db := database.DB

	book := entities.Book{}
	id := utils.ConvertToInt(c.Params("id"))

	r := db.Model(book).Where("id = ?", id).Take(&book)
	if r.Error != nil {
		return errors.SqlError(c, r.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       book,
	})
}

// @Tags    books
// @Summary Create a new book
// @Param   body      body     dto.CreateBookBody true " "
// @Success 201       {object} common.HttpResponse{data=entities.Book}
// @Router  /v1/books [post]
func (s BookService) Create(c *fiber.Ctx) error {
	db := database.DB

	body := dto.CreateBookBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	if body.UserID != nil {
		user := entities.User{}
		r := db.Model(user).Where("id = ?", body.UserID).Take(&user)
		if r.Error != nil {
			return errors.SqlError(c, r.Error)
		}
	}

	book := entities.Book{}
	copier.Copy(&book, &body)

	r := db.Create(&book)
	if r.Error != nil {
		return errors.SqlError(c, r.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Tags    books
// @Summary Update a book
// @Param   id             path     int true " "
// @Param   body           body     dto.UpdateBookBody true " "
// @Success 201            {object} common.HttpResponse{data=entities.Book}
// @Router  /v1/books/{id} [put]
func (s BookService) Update(c *fiber.Ctx) error {
	db := database.DB

	body := dto.UpdateBookBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	book := entities.Book{}
	id := utils.ConvertToInt(c.Params("id"))

	// q := db.Model(book).Session(&gorm.Session{})
	// r1 := q.Where("id = ?", id).Take(&book)
	r1 := db.Model(book).Where("id = ?", id).Take(&book)
	if r1.Error != nil {
		return errors.SqlError(c, r1.Error)
	}

	copier.CopyWithOption(&book, &body, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	// r2 := q.Where("id = ?", id).Updates(&book)
	r2 := db.Updates(&book)
	if r2.Error != nil {
		return errors.SqlError(c, r2.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Security BearerAuth
// @Summary  Delete a book
// @Tags     books
// @Param    id             path     int true " "
// @Success  200            {object} common.HttpResponse{}
// @Router   /v1/books/{id} [delete]
func (s BookService) Delete(c *fiber.Ctx) error {
	db := database.DB

	// user, err, ok := utils.CurrentUser(c)
	// if !ok {
	// 	return err
	// }

	book := entities.Book{}
	id := utils.ConvertToInt(c.Params("id"))

	r1 := db.Model(book).Where("id = ?", id).Take(&book)
	if r1.Error != nil {
		return errors.SqlError(c, r1.Error)
	}

	r2 := db.Delete(&book)
	if r2.Error != nil {
		return errors.SqlError(c, r2.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
	})
}
