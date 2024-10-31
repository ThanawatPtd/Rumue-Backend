package requests

type GetInsuranceRequest struct{
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year string `json:"year"`
}