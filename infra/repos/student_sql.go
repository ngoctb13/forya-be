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

func (s *studentSQLRepo) CreateStudent(ctx context.Context, student *models.Student) error {
	if err := s.db.WithContext(ctx).Create(student).Error; err != nil {
		return err
	}
	return nil
}

func (s *studentSQLRepo) BatchCreate(ctx context.Context, students []*models.Student) error {
	return s.db.WithContext(ctx).Create(students).Error
}

func (s *studentSQLRepo) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	if err := s.db.WithContext(ctx).Save(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentSQLRepo) DeleteStudentByID(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	if err := s.db.WithContext(ctx).First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := s.db.WithContext(ctx).
		Model(&student).
		Update("is_active", false).Error; err != nil {
		return nil, err
	}

	student.IsActive = false
	return &student, nil
}

func (s *studentSQLRepo) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	if err := s.db.WithContext(ctx).First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
