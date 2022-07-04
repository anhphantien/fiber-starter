package routers

import (
	"fiber-starter/enums"
	"fiber-starter/handlers"
	"fiber-starter/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(r fiber.Router) {
	r.Use(
		// middlewares.JwtAuth(),
		middlewares.RoleAuth(
			enums.User.Role.ADMIN,
			enums.User.Role.USER,
		),
	)

	r.Get("books", handlers.BookHandler{}.GetList)

	r.Get("books/:id", handlers.BookHandler{}.GetByID)

	r.Post("books", handlers.BookHandler{}.Create)

	r.Put("books/:id", handlers.BookHandler{}.Update)

	r.Delete("books/:id",
		middlewares.JwtAuth(),
		// middlewares.AdminRole,
		middlewares.UserRole,
		handlers.BookHandler{}.Delete,
	)
}
