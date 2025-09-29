package inputs

type ListCoursesInput struct {
	Name         *string
	Description  *string
	SessionCount *int
	PriceMax     *int
	PriceMin     *int
	OrderBy      *string
}

type CreateCourseInput struct {
	Name            string
	Description     string
	SessionCount    int
	PricePerSession int
}

type UpdateCourseInput struct {
	CourseID string
	Fields   map[string]interface{}
}
