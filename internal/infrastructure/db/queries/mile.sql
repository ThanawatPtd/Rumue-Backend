-- name: CreateMile :one
INSERT INTO "mile" (
    id, rate
) VALUES (
    $1, $2
)
RETURNING id, rate;

-- name: GetMiles :many
SELECT
    id, 
    rate
FROM "mile";

-- name: GetMile :one
SELECT
    rate
FROM "mile"
WHERE id=$1;


-- name: DeleteMile :exec
DELETE FROM "mile"
WHERE id=$1;