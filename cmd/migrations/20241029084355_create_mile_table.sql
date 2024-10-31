-- +goose Up
-- +goose StatementBegin
CREATE TABLE "mile" (
    id VARCHAR(100) PRIMARY KEY,
    rate FLOAT NOT NULL 
);

INSERT INTO "mile" (id, rate) VALUES
    ('boundary', 0.1),
    ('low', 0.2),      -- rate per 10,000 km/year
    ('mid', 0.3),      -- rate per 10,000 - 20,000 km/year
    ('high', 0.4),     -- rate per 20,000 km/year
    ('extra',0.5)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "mile";
-- +goose StatementEnd
