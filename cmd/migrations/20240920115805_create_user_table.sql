-- +goose Up
-- Create user table first
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
-- +goose Down
DROP TABLE post;
