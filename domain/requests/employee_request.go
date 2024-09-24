package requests

import "github.com/jackc/pgx/v5/pgtype"

type CreateEmployeeRequest struct {
	ID     pgtype.UUID   `json:"id"`
	Salary pgtype.Float4 `json:"salary"`
}