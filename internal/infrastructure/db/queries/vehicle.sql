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
RETURNING id, registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year, vehicle_color, vehicle_number, vehicle_number_location, engine_brand, engine_number, engine_number_location, fuel_type, chasis_number, wheel_type, total_piston, cc, horse_power, weight_unlanden, weight_laden, seating_capacity, miles;

