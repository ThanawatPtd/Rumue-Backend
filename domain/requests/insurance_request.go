package requests

type GetInsuranceRequest struct{
	Brand string `json:"brand" validate:"required"`
	Model string `json:"model" validate:"required"`
	Year string `json:"year" validate:"required"`
}