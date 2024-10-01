-- +goose Up
-- +goose StatementBegin
-- Vehicle Owner Table
CREATE TABLE "vehicle_owner"(
    user_id UUID,
    vehicle_id UUID,
    PRIMARY KEY (id, vehicle_id),
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES "vehicle"(vehicle_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
