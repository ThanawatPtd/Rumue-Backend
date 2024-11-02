-- name: GetAllTransactions :many
SELECT
    id,
    user_id,
    vehicle_id,
    employee_id,
    price,
    insurance_type,
    status,
    e_slip_image_url,
    cr_image_url,
    created_at,
    updated_at 
FROM "transaction";

-- name: CreateTransaction :one
INSERT INTO "transaction" (
    user_id, vehicle_id, price, insurance_type, status, e_slip_image_url, cr_image_url, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
)
RETURNING id, user_id, vehicle_id, price, insurance_type, status, e_slip_image_url, cr_image_url, cip_number,
vip_number;

-- name: GetAllTransactionsByUserID :many
SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.user_id = $1;

-- name: UpdateTransaction :exec
UPDATE "transaction"
SET
    employee_id = $2,
    status = $3,
    cip_number = $4,
    vip_number = $5,
    receipt_date = $6
WHERE id = $1;

-- name: GetTransactionByID :one
SELECT 
    id,
    user_id,
    vehicle_id,
    employee_id,
    price,
    insurance_type,
    status,
    e_slip_image_url,
    cr_image_url,
    receipt_date,
    created_at,
    updated_at 
FROM "transaction"
WHERE id = $1;

-- name: GetUserVehicleTransactionByID :one
SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.id = $1;


-- name: GetExpiredInsuranceTransactions :many

SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
    u.email, u.fname, u.lname, u.phone_number,
    u.address, u.nationality, u.birth_date, u.citizen_id,
    v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, 
    v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, 
    v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, 
    v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, 
    v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, 
    t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" AS t
JOIN "vehicle" AS v ON t.vehicle_id = v.id 
JOIN "user" AS u ON t.user_id = u.id
WHERE t.updated_at + INTERVAL '1 year' 
      BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '1 week';
