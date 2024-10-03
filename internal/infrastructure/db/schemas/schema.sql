CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) NOT NULL UNIQUE,
    fname VARCHAR(100) NOT NULL,
    lname VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);

-- EMPLOYEE Table (Inherits from USER)
CREATE TABLE "employee" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    salary REAL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
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
    engine_number VARCHAR(100) NOT NULL,
    chasis_number VARCHAR(100) NOT NULL,
    fuel_type VARCHAR(100) NOT NULL,
    horse_power INT NOT NULL,
    seating_capacity INT NOT NULL,
    weight_unlanden FLOAT NOT NULL,
    weight_laden FLOAT NOT NULL,
    tire_count INT NOT NULL,
    compulsory_insurance_policy_number VARCHAR(50) NOT NULL,
    voluntary_insurance_policy_number VARCHAR(50),
    insurance_type VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);

-- Vehicle Owner Table
CREATE TABLE "vehicle_owner"(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    vehicle_id UUID,
    UNIQUE (user_id, vehicle_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES "vehicle"(id) ON DELETE CASCADE
);

-- Transaction Table
CREATE TABLE "transaction" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vehicle_owner_id UUID NOT NULL,
    transaction_type VARCHAR(30) NOT NULL,
    transaction_status VARCHAR(100) NOT NULL,
    request_date TIMESTAMPTZ NOT NULL,
    response_date TIMESTAMPTZ,
    e_slip_image_url VARCHAR(100),
    FOREIGN KEY (vehicle_owner_id) REFERENCES "vehicle_owner"(id) ON DELETE CASCADE
);


-- Invoice Table
CREATE TABLE "invoice" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL,
    price FLOAT NOT NULL,
    invoice_image_url VARCHAR(100) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES "transaction"(id) ON DELETE CASCADE
);
