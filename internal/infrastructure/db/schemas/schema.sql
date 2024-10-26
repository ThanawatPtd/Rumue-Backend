CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) NOT NULL UNIQUE,
    fname VARCHAR(100) NOT NULL,
    lname VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    address VARCHAR(100) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    citizen_id VARCHAR(15) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- EMPLOYEE Table (Inherits from USER)
CREATE TABLE "employee" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    salary FLOAT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE
);

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

-- Vehicle Owner Table
CREATE TABLE "vehicle_owner"(
    user_id UUID,
    vehicle_id UUID,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (user_id, vehicle_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES "vehicle"(id) ON DELETE CASCADE
);

-- Transaction Table
CREATE TABLE "transaction" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    vehicle_id UUID,
    price FLOAT NOT NULL,
    insurance_type VARCHAR(30) NOT NULL,
    status VARCHAR(100) NOT NULL,
    e_slip_image_url VARCHAR(100) NOT NULL,
    cr_image_url VARCHAR(100) NOT NULL,
    cip_number VARCHAR(30),
    vip_number VARCHAR(30),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (user_id, vehicle_id) REFERENCES "vehicle_owner"(user_id, vehicle_id) ON DELETE CASCADE
);
