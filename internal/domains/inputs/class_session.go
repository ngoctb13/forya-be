package inputs

import "time"

type CreateClassSessionInput struct {
	ClassID string
	HeldAt  time.Time
}
