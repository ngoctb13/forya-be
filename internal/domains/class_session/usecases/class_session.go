package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	classRepo "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	"github.com/ngoctb13/forya-be/internal/domains/class_session/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type ClassSession struct {
	classSessionRepo        repos.IClassSession
	classRepo               classRepo.IClassRepo
	classSessionAttendances repos.IClassSessionAttendance
}

func NewClassSession(classSessionRepo repos.IClassSession, classRepo classRepo.IClassRepo, attendanceRepo repos.IClassSessionAttendance) *ClassSession {
	return &ClassSession{
		classSessionRepo:        classSessionRepo,
		classRepo:               classRepo,
		classSessionAttendances: attendanceRepo,
	}
}

func (cl *ClassSession) CreateClassSession(ctx context.Context, input *inputs.CreateClassSessionInput) error {
	session := &models.ClassSession{
		Name:    input.Name,
		ClassID: input.ClassID,
		HeldAt:  input.HeldAt,
	}

	return cl.classSessionRepo.Create(ctx, session)
}

func (cl *ClassSession) ListClassSessions(ctx context.Context, input *inputs.ListClassSessionsInput) ([]*models.ClassSession, *models.Pagination, error) {
	queries := make(map[string]interface{})
	if input.ClassID != nil {
		queries["class_id"] = *input.ClassID
	}
	if input.StartTime != nil {
		queries["start_time"] = *input.StartTime
	}
	if input.EndTime != nil {
		queries["end_time"] = *input.EndTime
	}

	pagination := models.NewPagination(input.Page, input.Limit)

	csArr, p, err := cl.classSessionRepo.List(ctx, queries, pagination)
	if err != nil {
		return nil, nil, err
	}

	classIDSet := make(map[string]bool)
	for _, cs := range csArr {
		if cs.ClassID != "" {
			classIDSet[cs.ClassID] = true
		}
	}

	classIDs := make([]string, 0, len(classIDSet))
	for id := range classIDSet {
		classIDs = append(classIDs, id)
	}

	if len(classIDs) > 0 {
		classMap, err := cl.classRepo.GetClassesByIDs(ctx, classIDs)
		if err != nil {
			return nil, nil, err
		}

		for _, cs := range csArr {
			if class, exists := classMap[cs.ClassID]; exists {
				cs.Class = class
			}
		}
	}

	return csArr, p, nil
}

func (cl *ClassSession) MarkAttendance(ctx context.Context, input *inputs.MarkClassSessionAttendanceInput) (*models.ClassSessionAttendance, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.ClassSessionID == "" {
		return nil, errors.New("class session id is required")
	}
	if input.CourseStudentID == "" {
		return nil, errors.New("course student id is required")
	}

	session, err := cl.classSessionRepo.GetByID(ctx, input.ClassSessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors.New("class session not found")
	}

	attendance, err := cl.classSessionAttendances.MarkAttendance(ctx, input.ClassSessionID, input.CourseStudentID, input.IsAttended)
	if err != nil {
		return nil, err
	}

	return attendance, nil
}
