package inputs

type CreateClassInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SearchClassByNameInput struct {
	Name  *string `json:"name"`
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
}
