-- name: GetAllUsers :many
SELECT *
FROM "user";

-- name: GetUserByID :one
SELECT *
FROM "user"
WHERE id = $1;

-- name: GetUserIDByEmail :one
SELECT
    id
FROM "user"
WHERE email = $1;

-- name: GetUserIDPasswordByEmail :one
SELECT
    id,
    password
FROM "user"
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO "user" (
    email, password, fname, lname, phone_number, address, nationality, birth_date, citizen_id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()
);


-- name: UpdateUser :one
UPDATE "user"
SET
    email = $2,
    fname = $3,
    lname = $4,
    phone_number = $5,
    address = $6,
    nationality = $7,
    birth_date = $8,
    citizen_id = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING id, email, fname, lname, phone_number, address, nationality, birth_date, citizen_id;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
