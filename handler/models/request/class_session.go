package request

type CreateClassSessionRequest struct {
	Name    string `json:"name"`
	ClassID string `json:"class_id"`
	HeldAt  string `json:"held_at"`
}

type ListClassSessionsRequest struct {
	ClassID   *string `form:"class_id"`
	StartTime *string `form:"start_time"`
	EndTime   *string `form:"end_time"`
	Page      int     `form:"page"`
	Limit     int     `form:"limit"`
}
