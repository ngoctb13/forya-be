package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type studentSQLRepo struct {
	db *gorm.DB
}

func NewStudentSQLRepo(db *gorm.DB) *studentSQLRepo {
	return &studentSQLRepo{
		db: db,
	}
}

func (s *studentSQLRepo) CreateStudent(ctx context.Context, class *models.Student) error {
	return nil
}

func (s *studentSQLRepo) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	return nil, nil
}

func (s *studentSQLRepo) DeleteStudentByID(ctx context.Context, id string) (*models.Student, error) {
	return nil, nil
}

func (s *studentSQLRepo) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	return nil, nil
}
