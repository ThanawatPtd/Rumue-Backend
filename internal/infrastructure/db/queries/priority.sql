-- name: CreatePriority :one
INSERT INTO "priority" (
    id, rate
) VALUES (
    $1, $2
)
RETURNING id, rate;

-- name: GetPriorities :many
SELECT
    id, 
    rate
FROM "priority";

-- name: GetPriority :one
SELECT
    rate
FROM "priority"
WHERE id = $1;

-- name: DeletePriority :exec
DELETE FROM "priority"
WHERE id = $1;
