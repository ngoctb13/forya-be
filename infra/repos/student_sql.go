package repos

import (
	"context"
	"time"

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

func (s *studentSQLRepo) GetStudentsByClassID(ctx context.Context, classID string, queryOpts models.QueryOptions) ([]*models.ClassEnrollments, error) {
	var results []struct {
		models.Student
		JoinedAt time.Time  `gorm:"column:joined_at"`
		LeftAt   *time.Time `gorm:"column:left_at"`
	}

	q := s.db.WithContext(ctx).
		Table("students").
		Select("students.*, cs.joined_at, cs.left_at").
		Joins("JOIN class_student cs ON cs.student_id = students.id").
		Where("cs.class_id = ?", classID)

	if queryOpts.JoinedAt != nil {
		q = q.Where("cs.joined_at >= ?", queryOpts.JoinedAt)
	}

	if queryOpts.LeftAt != nil {
		q = q.Where("cs.left_at >= ?", queryOpts.LeftAt)
	}

	if err := q.Find(&results).Error; err != nil {
		return nil, err
	}

	var enriched []*models.ClassEnrollments
	for _, r := range results {
		enriched = append(enriched, &models.ClassEnrollments{
			Student:  r.Student,
			JoinedAt: r.JoinedAt,
			LeftAt:   r.LeftAt,
		})
	}

	return enriched, nil
}
