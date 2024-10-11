-- +goose Up
-- +goose StatementBegin
-- Create transaction table
CREATE TABLE "transaction" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    vehicle_id UUID,
    insurance_type VARCHAR(30) NOT NULL,
    transaction_status VARCHAR(100) NOT NULL,
    request_date TIMESTAMPTZ NOT NULL,
    response_date TIMESTAMPTZ,
    e_slip_image_url VARCHAR(100) NOT NULL,
    car_registration_image_url VARCHAR(100) NOT NULL,
    compulsory_insurance_policy_number VARCHAR(30),
    voluntary_insurance_policy_number VARCHAR(30),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (user_id, vehicle_id) REFERENCES "vehicle_owner"(user_id, vehicle_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "transaction";
-- +goose StatementEnd
