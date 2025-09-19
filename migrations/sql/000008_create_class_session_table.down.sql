DROP INDEX IF EXISTS idx_attendance_student_course;
DROP INDEX IF EXISTS idx_attendance_session;
DROP INDEX IF EXISTS idx_class_sessions_class;
DROP INDEX IF EXISTS idx_student_courses_course;
DROP INDEX IF EXISTS idx_student_courses_student;

DROP TABLE IF EXISTS class_session_attendance;
DROP TABLE IF EXISTS class_session;
