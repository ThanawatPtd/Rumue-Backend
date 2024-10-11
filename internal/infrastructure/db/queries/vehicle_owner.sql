-- name: GetAllVehicleOwner :many
SELECT
    user_id,
    vehicle_id
from "vehicle_owner";

-- name: GetAllVehicleOwnerByUserId :many
SELECT
    user_id,
    vehicle_id
from "vehicle_owner"
WHERE user_id = $1;

-- name: GetVehicleOwnerByBothId :one
SELECT
    user_id,
    vehicle_id
from "vehicle_owner"
WHERE user_id = $1 and vehicle_id = $2;

-- name: CreateVehicleOwner :one
INSERT INTO "vehicle_owner"(
    user_id, vehicle_id, created_at, updated_at
) VALUES(
    $1, $2, NOW(), NOW()
)RETURNING user_id, vehicle_id;
