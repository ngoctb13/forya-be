package outputs

import (
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ListClassSessionsOutput struct {
	ID     string
	Name   string
	HeldAt time.Time
	Class  struct {
		ID   string
		Name string
	}
}

func ToListClassSessionsOutput(cs []*models.ClassSession) []*ListClassSessionsOutput {
	var outs []*ListClassSessionsOutput
	for _, c := range cs {
		out := &ListClassSessionsOutput{
			ID:     c.ID,
			Name:   c.Name,
			HeldAt: c.HeldAt,
		}
		if c.Class != nil {
			out.Class.ID = c.Class.ID
			out.Class.Name = c.Class.Name
		}

		outs = append(outs, out)
	}

	return outs
}
