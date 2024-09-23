-- name: GetAllAdmins :many
SELECT
    id,
    created_at,
    updated_at
FROM "admin";

-- name: GetAdminByID :one
SELECT
    id,
    created_at,
    updated_at
FROM "admin"
WHERE id = $1;

-- name: CreateAdmin :one
INSERT INTO "admin" (
    id, created_at, updated_at
) VALUES (
    $1, NOW(), NOW()
)
RETURNING id, created_at;

-- name: UpdateAdmin :one
UPDATE "admin"
SET
    updated_at = NOW()
WHERE id = $1
RETURNING id, updated_at;

-- name: DeleteAdmin :exec
DELETE FROM "admin"
WHERE id = $1;
