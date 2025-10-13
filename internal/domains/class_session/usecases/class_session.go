package usecases

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class_session/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
)

type ClassSession struct {
	classSessionRepo repos.IClassSession
}

func NewClassSession(classSessionRepo repos.IClassSession) *ClassSession {
	return &ClassSession{
		classSessionRepo: classSessionRepo,
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

func (cl *ClassSession) ListClassSessions(ctx context.Context, input *inputs.ListClassSessionsInput) ([]*outputs.ListClassSessionsOutput, *models.Pagination, error) {
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

	return outputs.ToListClassSessionsOutput(csArr), p, err
}
