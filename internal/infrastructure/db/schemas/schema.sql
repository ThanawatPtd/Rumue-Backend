CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    fname VARCHAR(100) NOT NULL,
    lname VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(10) NOT NULL,
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);

ALTER TABLE "user" ALTER COLUMN id SET DEFAULT uuid_generate_v4();

-- EMPLOYEE Table (Inherits from USER)
CREATE TABLE "employee" (
    id UUID PRIMARY KEY,
    salary REAL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE
);

ALTER TABLE "employee" ALTER COLUMN id SET DEFAULT uuid_generate_v4();

-- ADMIN Table (Inherits from USER)
CREATE TABLE "admin" (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE
);

ALTER TABLE "admin" ALTER COLUMN id SET DEFAULT uuid_generate_v4();

-- EMPLOYEE_MANAGEMENT Table
CREATE TABLE "employee_management" (
    employee_id UUID,
    admin_id UUID,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    PRIMARY KEY (employee_id, admin_id),
    FOREIGN KEY (employee_id) REFERENCES "employee"(id) ON DELETE CASCADE,
    FOREIGN KEY (admin_id) REFERENCES "admin"(id) ON DELETE CASCADE
);

ALTER TABLE "employee_management" ALTER COLUMN employee_id SET DEFAULT uuid_generate_v4();
ALTER TABLE "employee_management" ALTER COLUMN admin_id SET DEFAULT uuid_generate_v4();


-- Vehicle Table
CREATE TABLE "vehicle"(
    vehicle_id UUID PRIMARY KEY,
    registration_date TIMESTAMPTZ NOT NULL,
    registration_number TIMESTAMPTZ NOT NULL,
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

ALTER TABLE "vehicle" ALTER COLUMN vehicle_id SET DEFAULT uuid_generate_v4();

-- Vehicle Owner Table
CREATE TABLE "vehicle_owner"(
    user_id UUID,
    vehicle_id UUID,
    PRIMARY KEY (user_id, vehicle_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES "vehicle"(vehicle_id) ON DELETE CASCADE
);
