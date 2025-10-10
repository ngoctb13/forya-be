package repos

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

func ApplyPagination(db *gorm.DB, page, limit int) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	return db.Offset(offset).Limit(limit)
}

func Paginate[T any](db *gorm.DB, page, limit int, out *[]T) (*models.Pagination, error) {
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := ApplyPagination(db, page, limit).Find(out).Error; err != nil {
		return nil, err
	}

	return models.NewPagination(page, limit, int(total)), nil
}
