-- name: GetAllEmployees :many
SELECT
    e.id,
    e.fname AS name,
    e.email,
    e.salary,
    e.position,
    e.created_at,
    e.updated_at
FROM employee e;

-- name: GetEmployeeByID :one
SELECT
    e.id,
    e.fname AS name,
    e.email,
    e.salary,
    e.position,
    e.created_at,
    e.updated_at
FROM employee e
WHERE e.id = $1;

-- name: CreateEmployee :one
INSERT INTO employee (
    id, email, fname, lname, password, phone_number, address, salary, position, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()
)
RETURNING id, fname AS name, email, salary, position, created_at, updated_at;

-- name: UpdateEmployee :one
UPDATE employee
SET
    email = $2,
    fname = $3,
    lname = $4,
    password = $5,
    phone_number = $6,
    address = $7,
    salary = $8,
    position = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING id, fname AS name, email, salary, position, created_at, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = $1;
