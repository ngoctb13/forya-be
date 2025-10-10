package models

type Pagination struct {
	Page      int
	Limit     int
	Total     int
	TotalPage int
}

func NewPagination(page, limit, total int) *Pagination {
	if limit <= 0 {
		limit = 10
	}

	totalPage := (total + limit - 1) / limit

	return &Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
	}
}

func (p *Pagination) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit
}
