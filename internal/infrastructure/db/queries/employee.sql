-- name: GetAllEmployees :many
SELECT
   id,
   salary,
   created_at,
   updated_at 
FROM "employee";

-- name: GetEmployeeByID :one
SELECT
    id,
    salary,
    created_at,
    updated_at 
FROM "employee"
WHERE id = $1;

-- name: CreateEmployee :one
INSERT INTO "employee" (
    id, salary, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING id, salary, created_at;

-- name: UpdateEmployee :one
UPDATE "employee"
SET
    salary = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, salary, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM "employee"
WHERE id = $1;
