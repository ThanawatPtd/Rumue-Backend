package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type AdminRepository interface {
	ListAll(c *context.Context) (*[]dbmodel.Admin, error)
	Save(c *context.Context, id *pgtype.UUID) (*dbmodel.Admin, error)
	GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.Admin, error)
	Update(c *context.Context, id *pgtype.UUID) (*dbmodel.Admin, error)
	Delete(c *context.Context, id *pgtype.UUID) (error)
}