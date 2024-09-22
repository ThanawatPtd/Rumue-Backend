-- name: GetAllEmployees :many
SELECT
    *
FROM "employee";

-- name: GetEmployeeByID :one
SELECT
    *
FROM "employee"
WHERE id = $1;

-- name: CreateEmployee :one
INSERT INTO "employee" (
    id, salary, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING *;

-- name: UpdateEmployee :one
UPDATE "employee"
SET
    salary = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM "employee"
WHERE id = $1;
