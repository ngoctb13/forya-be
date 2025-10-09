package request

type CreateClassSessionRequest struct {
	Name    string `json:"name"`
	ClassID string `json:"class_id"`
	HeldAt  string `json:"held_at"`
}
