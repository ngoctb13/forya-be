package models

type CreateClassRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
