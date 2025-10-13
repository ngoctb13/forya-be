package models

import "gorm.io/gorm"

type Pagination struct {
	Page      int
	Limit     int
	Total     int64
	TotalPage int
}

func NewPagination(page, limit int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p *Pagination) ApplyToQuery(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return db.Limit(p.Limit).Offset(offset)
}

func (p *Pagination) SetTotal(total int64) {
	p.Total = total
	p.TotalPage = int((total + int64(p.Limit) - 1) / int64(p.Limit))
}
