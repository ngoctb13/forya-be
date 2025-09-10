CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE classes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    schedule VARCHAR(100),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);