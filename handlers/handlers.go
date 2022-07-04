package handlers

import "fiber-starter/repositories"

var (
	bookRepository = repositories.BookRepository{}
	userRepository = repositories.UserRepository{}
)
