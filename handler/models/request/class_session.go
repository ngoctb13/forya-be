package request

type CreateClassSessionRequest struct {
	ClassID string `json:"class_id"`
	HeldAt  string `json:"held_at"`
}
