package entities

type VehicleOwner struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	VehicleID string `json:"vehicleId"`	
}