-- name: GetAllVehicleOwner :many
SELECT
    user_id,
    vehicle_id
from "vehicle_owner";

-- name: GetAllVehicleOwnerByUserId :many
SELECT
    vehicle_id
from "vehicle_owner"
WHERE user_id = $1;

-- name: CreateVehicleOwner :one
INSERT INTO "vehicle_owner"(
    user_id,vehicle_id
) VALUES(
    $1, $2
)RETURNING user_id, vehicle_id;
