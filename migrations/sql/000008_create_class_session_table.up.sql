CREATE TABLE class_session (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    class_id UUID NOT NULL,
    held_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_class FOREIGN KEY (class_id) REFERENCES classes(id)
);

CREATE TABLE class_session_attendance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    class_session_id UUID NOT NULL,
    course_student_id UUID NOT NULL,
    is_attended BOOLEAN DEFAULT FALSE,

    CONSTRAINT fk_session FOREIGN KEY (class_session_id) REFERENCES class_session(id),
    CONSTRAINT fk_course_student FOREIGN KEY (course_student_id) REFERENCES course_student(id)
);

CREATE INDEX idx_student_courses_student ON course_student(student_id);
CREATE INDEX idx_student_courses_course ON course_student(course_id);
CREATE INDEX idx_class_sessions_class ON class_session(class_id);
CREATE INDEX idx_attendance_session ON class_session_attendance(class_session_id);
CREATE INDEX idx_attendance_student_course ON class_session_attendance(course_student_id);