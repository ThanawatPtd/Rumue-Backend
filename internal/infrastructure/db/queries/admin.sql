-- name: GetAllAdmins :many
SELECT
    *
FROM "admin";

-- name: GetAdminByID :one
SELECT
    *
FROM "admin"
WHERE id = $1;

-- name: CreateAdmin :one
INSERT INTO "admin" (
    id, created_at, updated_at
) VALUES (
    $1, NOW(), NOW()
)
RETURNING *;

-- name: UpdateAdmin :one
UPDATE "admin"
SET
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM "admin"
WHERE id = $1;
