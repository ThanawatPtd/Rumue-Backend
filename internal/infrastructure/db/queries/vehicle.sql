-- name: GetAllVehicle :many
SELECT
    id, registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year, vehicle_color, vehicle_number, vehicle_number_location, engine_brand, engine_number, engine_number_location, fuel_type, chasis_number, wheel_type, total_piston, cc, horse_power, weight_unlanden, weight_laden, seating_capacity, miles
FROM "vehicle"
WHERE id = $1;


-- name: CreateVehicle :one
INSERT INTO "vehicle" (
    registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year, vehicle_color, vehicle_number, vehicle_number_location, engine_brand, engine_number, engine_number_location, fuel_type, chasis_number, wheel_type, total_piston, cc, horse_power, weight_unlanden, weight_laden, seating_capacity, miles, created_at, updated_at
) VALUES (
   $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, NOW(), NOW()
)
RETURNING id;

-- name: FindAllTemplate :many
WITH latest_vehicles AS (
    SELECT
        v.vehicle_number,
        MAX(v.updated_at) AS latest_update
    FROM vehicle_owner vo
    JOIN vehicle v ON vo.vehicle_id = v.id
    WHERE vo.user_id = $1
    GROUP BY v.vehicle_number
)
SELECT
    v.id, v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category,
    v.characteristics, v.brand, v.model, v.model_year, 
    v.vehicle_color, v.vehicle_number,
    v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location,
    v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power,
    v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, v.updated_at
FROM vehicle_owner vo
JOIN vehicle v ON vo.vehicle_id = v.id
JOIN latest_vehicles lv ON v.vehicle_number = lv.vehicle_number AND v.updated_at = lv.latest_update
WHERE vo.user_id = $1;

