package inputs

import "time"

type CreateClassSessionInput struct {
	Name    string
	ClassID string
	HeldAt  time.Time
}
