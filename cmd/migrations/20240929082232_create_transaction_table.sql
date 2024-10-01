-- +goose Up
-- +goose StatementBegin
-- Create transaction table
CREATE TABLE "transaction" (
    transaction_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vehicle_id UUID,
    owner_id UUID,
    transaction_type VARCHAR(30) NOT NULL,
    transaction_status VARCHAR(100) NOT NULL,
    request_date TIMESTAMPTZ NOT NULL,
    response_date TIMESTAMPTZ,
    e_slip_image_url VARCHAR(100),
    FOREIGN KEY (vehicle_id, owner_id) REFERENCES "vehicle_owner"(vehicle_id, id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
