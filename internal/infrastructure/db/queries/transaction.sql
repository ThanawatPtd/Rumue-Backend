-- name: GetAllTransactions :many
SELECT
    id,
    vehicle_owner_id,
    transaction_type,
    transaction_status,
    request_date,
    response_date,
    e_slip_image_url
FROM "transaction";

-- name: CreateTransaction :one
INSERT INTO "transaction" (
    vehicle_owner_id, transaction_type, transaction_status, request_date, response_date, e_slip_image_url, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, NOW(), NOW()
)
RETURNING id, vehicle_owner_id, transaction_type, transaction_status, request_date, response_date, e_slip_image_url; 