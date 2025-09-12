CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE class_student (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    class_id UUID NOT NULL,
    student_id UUID NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    left_at TIMESTAMPTZ,

    CONSTRAINT fk_class FOREIGN KEY (class_id) REFERENCES classes(id),
    CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES students(id)
);
