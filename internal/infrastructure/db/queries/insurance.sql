-- name: CreateInsurance :one
INSERT INTO "insurance" (
    year, model, brand, price
) VALUES (
    $1, $2, $3, $4
)
RETURNING year, model, brand, price;

-- name: GetInsurances :many
SELECT
    model,
    brand,
    year
FROM "insurance"
ORDER BY brand, model, year;

-- name: GetInsurance :one
SELECT
    price
FROM "insurance"
WHERE brand = $1 AND model = $2 AND year = $3;

-- name: DeleteInsurance :exec
DELETE FROM "insurance"
WHERE brand = $1 AND model = $2 AND year = $3;
