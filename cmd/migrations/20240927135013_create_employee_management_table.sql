-- +goose Up
-- +goose StatementBegin
-- EMPLOYEE_MANAGEMENT Table
CREATE TABLE "employee_management" (
    employee_id UUID,
    admin_id UUID,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    PRIMARY KEY (employee_id, admin_id),
    FOREIGN KEY (employee_id) REFERENCES "employee"(id) ON DELETE CASCADE,
    FOREIGN KEY (admin_id) REFERENCES "admin"(id) ON DELETE CASCADE
);

ALTER TABLE "employee_management" ALTER COLUMN employee_id SET DEFAULT uuid_generate_v4();
ALTER TABLE "employee_management" ALTER COLUMN admin_id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
