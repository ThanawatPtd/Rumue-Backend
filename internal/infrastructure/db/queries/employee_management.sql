-- name: GetAllEmployeeManagement :many
SELECT
    employee_id,
    admin_id,
    created_at,
    updated_at
FROM "employee_management";

-- name: GetEmployeeManagementsByEmployeeID :many
SELECT
   employee_id,
   admin_id,
   created_at,
   updated_at 
FROM "employee_management"
WHERE employee_id = $1;

-- name: GetAllEmployeeManagementsByAdminID :many
SELECT
   employee_id,
   admin_id,
   created_at,
   updated_at 
FROM "employee_management"
WHERE admin_id = $1;

-- name: CreateEmployeeManagement :one
INSERT INTO "employee_management" (
    employee_id, admin_id, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING employee_id, admin_id, created_at;

-- name: UpdateEmployeeManagement :one
UPDATE "employee_management"
SET
    updated_at = NOW()
WHERE employee_id = $1 AND admin_id = $2
RETURNING employee_id, admin_id, updated_at;

-- name: DeleteEmployeeManagement :exec
DELETE FROM "employee_management"
WHERE employee_id = $1 AND admin_id = $2;