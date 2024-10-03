-- +goose Up
-- +goose StatementBegin
-- Vehicle Owner Table
CREATE TABLE "vehicle_owner"(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    vehicle_id UUID,
    UNIQUE (user_id, vehicle_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES "vehicle"(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
