-- name: GetAllVehicle :many
SELECT
    id,
    registration_date,
    registration_number,
    province,
    vehicle_type ,
    vehicle_category,
    characteristics ,
    brand,
    model ,
    model_year ,
    vehicle_color ,
    engine_number ,
    chasis_number ,
    fuel_type ,
    horse_power,
    seating_capacity,
    weight_unlanden ,
    weight_laden ,
    tire_count,
    compulsory_insurance_policy_number ,
    voluntary_insurance_policy_number ,
    insurance_type ,
    created_at,
    updated_at
from "vehicle";

-- name: CreateVehicle :one
INSERT INTO "vehicle" (
    registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year,vehicle_color, engine_number, chasis_number , fuel_type, horse_power, seating_capacity, weight_unlanden, weight_laden, tire_count, compulsory_insurance_policy_number, voluntary_insurance_policy_number, insurance_type, created_at, updated_at
) VALUES (
   $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, NOW(), NOW()
)
RETURNING id,registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year,vehicle_color, engine_number, chasis_number , fuel_type, horse_power, seating_capacity, weight_unlanden, weight_laden, tire_count, compulsory_insurance_policy_number, voluntary_insurance_policy_number, insurance_type, created_at;
