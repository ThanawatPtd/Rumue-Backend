CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
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