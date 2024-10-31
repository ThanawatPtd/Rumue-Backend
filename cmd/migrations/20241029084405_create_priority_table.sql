-- +goose Up
-- +goose StatementBegin
CREATE TABLE "priority" (
    id VARCHAR(100) PRIMARY KEY,
    rate FLOAT NOT NULL 
);
INSERT INTO "priority" (id, rate) VALUES
    (0, 1.0),
    (1, 2.0),
    (2, 1.5),
    (3, 1.2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "priority";
-- +goose StatementEnd
