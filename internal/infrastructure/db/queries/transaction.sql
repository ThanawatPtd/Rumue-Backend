-- name: GetAllTransactions :many
SELECT
    id,
    user_id,
    vehicle_id,
    insurance_type,
    transaction_status,
    request_date,
    response_date,
    e_slip_image_url,
    car_registration_image_url,
    compulsory_insurance_policy_number,
    voluntary_insurance_policy_number 
FROM "transaction";

-- name: CreateTransaction :one
INSERT INTO "transaction" (
    user_id, vehicle_id, insurance_type, transaction_status, request_date, e_slip_image_url, car_registration_image_url, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, NOW(), $5, $6, NOW(), NOW()
)
RETURNING id, user_id, vehicle_id, insurance_type, transaction_status, request_date, response_date, e_slip_image_url, car_registration_image_url, compulsory_insurance_policy_number,
voluntary_insurance_policy_number;
