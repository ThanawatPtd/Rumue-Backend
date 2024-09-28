-- +goose Up
-- +goose StatementBegin
-- ADMIN Table (Inherits from USER)
CREATE TABLE "admin" (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE
);

ALTER TABLE "admin" ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
