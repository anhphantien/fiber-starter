package dto

type CreateBookBody struct {
	Title       *string `json:"title" validate:"omitempty,max=255" example:"abc"`
	Description *string `json:"description" validate:"omitempty,max=255" example:"abc"`
	Content     *string `json:"content" validate:"omitempty" example:"abc"`
}
