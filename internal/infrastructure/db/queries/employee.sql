-- name: GetAllEmployees :many
SELECT
   id,
   salary,
   created_at,
   updated_at 
FROM "employee";

-- name: GetEmployeeByID :one
SELECT
    id,
    salary
FROM "employee"
WHERE id = $1;

-- name: CreateEmployee :one
INSERT INTO "employee" (
    id, salary, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING id, salary, created_at;

-- name: UpdateEmployee :one
UPDATE "employee"
SET
    salary = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, salary, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM "employee"
WHERE id = $1;

-- name: FindInsuranceToday :many
SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.status = 'pending'
or t.updated_at::date = CURRENT_DATE;
