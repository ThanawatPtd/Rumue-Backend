-- name: GetAllUsers :many
SELECT
    id,
    email,
    fname,
    lname,
    phone_number,
    password,
    address,
    created_at,
    updated_at
FROM "user";

-- name: GetUserByID :one
SELECT
    id,
    email,
    fname,
    lname,
    phone_number,
    password,
    address
FROM "user"
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    fname,
    lname,
    phone_number,
    password,
    address
FROM "user"
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO "user" (
    email, fname, lname, password, phone_number, address, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, NOW(), NOW()
)
RETURNING id, email, fname, lname, phone_number, address;

-- name: UpdateUser :one
UPDATE "user"
SET
    email = $2,
    fname = $3,
    lname = $4,
    phone_number = $5,
    address = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING id, email, fname, lname, phone_number, address;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
