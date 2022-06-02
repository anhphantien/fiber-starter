package dto

type Sort struct {
	Field string
	Order string
}

type Pagination struct {
	Limit   int
	Page    int
	Keyword string
	Filter  map[string]any
	Sort    Sort
}
