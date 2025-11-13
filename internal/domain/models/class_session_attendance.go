package models

type ClassSessionAttendance struct {
	ID              string `gorm:"default:uuid_generate_v4()"`
	ClassSessionID  string `gorm:"column:class_session_id"`
	CourseStudentID string `gorm:"column:course_student_id"`
	IsAttended      bool   `gorm:"column:is_attended"`
}

func (ClassSessionAttendance) TableName() string {
	return "class_session_attendance"
}
