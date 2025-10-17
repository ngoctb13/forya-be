package inputs

type ListCoursesInput struct {
	Name         *string
	SessionCount *int
	PriceMax     *float64
	PriceMin     *float64
	OrderBy      *string
	Page         int
	Limit        int
}

type CreateCourseInput struct {
	Name            string
	Description     string
	SessionCount    int
	PricePerSession float64
}

type UpdateCourseInput struct {
	CourseID string
	Fields   UpdateCourseFields
}

type UpdateCourseFields struct {
	Name            *string
	Description     *string
	SessionCount    *int
	PricePerSession *float64
}
