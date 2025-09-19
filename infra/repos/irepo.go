package repos

import (
	classRp "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	classStudentRp "github.com/ngoctb13/forya-be/internal/domains/class_student/repos"
	courseRp "github.com/ngoctb13/forya-be/internal/domains/course/repos"
	courseStudentRp "github.com/ngoctb13/forya-be/internal/domains/course_student/repos"
	studentRp "github.com/ngoctb13/forya-be/internal/domains/student/repos"
	userRp "github.com/ngoctb13/forya-be/internal/domains/user/repos"
)

type IRepo interface {
	Users() userRp.IUserRepo
	Classes() classRp.IClassRepo
	Students() studentRp.IStudentRepo
	ClassStudent() classStudentRp.IClassStudentRepo
	Courses() courseRp.ICourseRepo
	CourseStudent() courseStudentRp.ICourseStudentRepo
}
