CREATE TABLE supplies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    unit VARCHAR(50) NOT NULL,                   -- đơn vị tính: cái, hộp, tờ,...
    min_threshold INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE supply_batches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supply_id UUID NOT NULL REFERENCES supplies(id),
    quantity INT NOT NULL,              -- tổng số nhập
    remaining_quantity INT NOT NULL,    -- còn lại trong kho
    purchase_price NUMERIC(12,2) NOT NULL,     -- giá nhập
    purchase_date TIMESTAMPTZ NOT NULL,
    contact VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_supply_batches_supply_id ON supply_batches (supply_id);
CREATE INDEX idx_supply_batches_purchase_date ON supply_batches (purchase_date);


CREATE TABLE supply_usages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    batch_id UUID NOT NULL REFERENCES supply_batches(id),
    student_id UUID NOT NULL,
    class_session_id UUID NOT NULL,
    quantity INT NOT NULL,
    unit_price NUMERIC(12,2) NOT NULL, -- giá bán cho học sinh (có thể khác giá nhập)
    total_price NUMERIC(12,2) NOT NULL,
    used_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_supply_usages_batch_id ON supply_usages (batch_id);
CREATE INDEX idx_supply_usages_student_id ON supply_usages (student_id);
CREATE INDEX idx_supply_usages_class_session_id ON supply_usages (class_session_id);
CREATE INDEX idx_supply_usages_used_at ON supply_usages (used_at);
