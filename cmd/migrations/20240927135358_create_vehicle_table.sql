-- +goose Up
-- +goose StatementBegin
-- Vehicle Table
CREATE TABLE "vehicle"(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    registration_date TIMESTAMPTZ NOT NULL,
    registration_number VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL,
    vehicle_type VARCHAR(100) NOT NULL,
    vehicle_category VARCHAR(100) NOT NULL,
    characteristics VARCHAR(100) NOT NULL,
    brand VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    model_year VARCHAR(100) NOT NULL,
    vehicle_color VARCHAR(50) NOT NULL,
    vehicle_number VARCHAR(50) NOT NULL,
    vehicle_number_location VARCHAR(100) NOT NULL,
    engine_brand VARCHAR(100) NOT NULL,
    engine_number VARCHAR(100) NOT NULL,
    engine_number_location VARCHAR(100) NOT NULL,
    chasis_number VARCHAR(100) NOT NULL,
    fuel_type VARCHAR(100) NOT NULL,
    wheel_type VARCHAR(100) NOT NULL,
    total_piston INT NOT NULL,
    cc INT NOT NULL,
    horse_power INT NOT NULL,
    seating_capacity INT NOT NULL,
    weight_unlanden FLOAT NOT NULL,
    weight_laden FLOAT NOT NULL,
    miles FLOAT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "vehicle";
-- +goose StatementEnd
