-- name: GetAllVehicleOwner :many
SELECT
    id,
    user_id,
    vehicle_id
from "vehicle_owner";

-- name: GetAllVehicleOwnerByUserId :many
SELECT
    id,
    user_id,
    vehicle_id
from "vehicle_owner"
WHERE user_id = $1;

-- name: CreateVehicleOwner :one
INSERT INTO "vehicle_owner"(
    user_id,vehicle_id
) VALUES(
    $1, $2
)RETURNING id, user_id, vehicle_id;
