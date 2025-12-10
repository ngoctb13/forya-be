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

func (s *studentSQLRepo) GetStudentsByClassID(ctx context.Context, classID string, queries map[string]interface{}, pagination *models.Pagination) ([]*models.ClassEnrollments, *models.Pagination, error) {
	var results []struct {
		models.Student
		JoinedAt time.Time  `gorm:"column:joined_at"`
		LeftAt   *time.Time `gorm:"column:left_at"`
	}

	query := s.db.WithContext(ctx).
		Table("students").
		Select("students.*, cs.joined_at, cs.left_at").
		Joins("JOIN class_student cs ON cs.student_id = students.id").
		Where("cs.class_id = ?", classID)

	for k, v := range queries {
		switch k {
		case "joined_at":
			query = query.Where("joined_at >= ?", v)
		case "left_at":
			query = query.Where("left_at >= ?", v)
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)
	query = pagination.ApplyToQuery(query)

	if err := query.Find(&results).Error; err != nil {
		return nil, nil, err
	}

	var enriched []*models.ClassEnrollments
	for _, r := range results {
		enriched = append(enriched, &models.ClassEnrollments{
			Student:  r.Student,
			JoinedAt: r.JoinedAt,
			LeftAt:   r.LeftAt,
		})
	}

	return enriched, pagination, nil
}

func (s *studentSQLRepo) UpdateWithMap(ctx context.Context, studentID string, fields map[string]interface{}) (*models.Student, error) {
	student := &models.Student{}

	if err := s.db.WithContext(ctx).
		Model(student).
		Where("id = ?", studentID).
		Updates(fields).Error; err != nil {
		return nil, err
	}

	if err := s.db.WithContext(ctx).
		First(student, "id = ?", studentID).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentSQLRepo) List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.Student, *models.Pagination, error) {
	query := s.db.WithContext(ctx).Model(&models.Student{})

	for k, v := range queries {
		switch k {
		case "full_name":
			if name, ok := v.(string); ok {
				query = query.Where("unaccent(lower(full_name)) ILIKE unaccent(lower(?))", "%"+name+"%")
			}
		case "age_min":
			query = query.Where("age >= ?", v)
		case "age_max":
			query = query.Where("age <= ?", v)
		case "phone_number":
			if pn, ok := v.(string); ok {
				query = query.Where("phone_number ILIKE ?", "%"+pn+"%")
			}
		case "parent_phone_number":
			if ppn, ok := v.(string); ok {
				query = query.Where("parent_phone_number ILIKE ?", "%"+ppn+"%")
			}
		}
	}

	var (
		total    int64
		students []*models.Student
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)
	query = pagination.ApplyToQuery(query)
	query = query.Order("created_at DESC")

	if err := query.Find(&students).Error; err != nil {
		return nil, nil, err
	}

	return students, pagination, nil
}
