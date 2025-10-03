package usecases

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class_session/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
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
		ClassID: input.ClassID,
		HeldAt:  input.HeldAt,
	}

	return cl.classSessionRepo.Create(ctx, session)
}
