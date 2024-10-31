-- name: GetAllUsers :many
SELECT 
    email,
    fname,
    lname,
    phone_number,
    address,
    nationality,
    birth_date,
    citizen_id
FROM "user";

-- name: GetUserProfile :one
SELECT 
    u.email,
    u.fname,
    u.lname,
    u.phone_number,
    u.address,
    u.nationality,
    u.birth_date,
    u.citizen_id,
    e.salary
FROM "user" u
LEFT JOIN employee e ON u.id = e.id
WHERE u.id = $1;

-- name: GetUserByID :one
SELECT
    email,
    fname,
    lname,
    phone_number,
    address,
    nationality,
    birth_date,
    citizen_id
FROM "user"
WHERE id = $1;

-- name: GetUserIDPasswordByEmail :one
SELECT
    id,
    password
FROM "user"
WHERE email = $1;

-- name: GetUserIDPasswordByID :one
SELECT
    id,
    password
FROM "user"
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO "user" (
    email, password, fname, lname, phone_number, address, nationality, birth_date, citizen_id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()
);

-- name: UpdateUser :one
UPDATE "user"
SET
    fname = $2,
    lname = $3,
    phone_number = $4,
    address = $5,
    nationality = $6,
    birth_date = $7,
    citizen_id = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING email, fname, lname, phone_number, address, nationality, birth_date, citizen_id;

-- name: UpdateUserPassword :exec
UPDATE "user"
SET
    password = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
