package dto

type Pagination struct {
	Limit   int
	Page    int
	Keyword string
	Filter  any
	Sort    any
}

// filter any { [key: string]: any } = {};
// sort: Sort = { field: 'id', order: 'DESC' };
