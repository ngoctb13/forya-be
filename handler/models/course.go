package models

type CreateCourseRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	SessionCount    int    `json:"session_count"`
	PricePerSession int    `json:"price_per_session"`
}
