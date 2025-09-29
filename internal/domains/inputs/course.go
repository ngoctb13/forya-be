package inputs

type ListCoursesInput struct {
	Name         *string
	Description  *string
	SessionCount *int
	PriceMax     *float64
	PriceMin     *float64
	OrderBy      *string
}

type CreateCourseInput struct {
	Name            string
	Description     string
	SessionCount    int
	PricePerSession float64
}

type UpdateCourseInput struct {
	CourseID string
	Fields   map[string]interface{}
}
