package psql

import (
	"context"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAdminRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresAdminRepository(db *pgxpool.Pool) repositories.AdminRepository {
	return &PostgresAdminRepository{
		Queries: dbmodel.New(db),
		DB: db,
	}
}

func (a *PostgresAdminRepository) ListAll(c *context.Context) (*[]dbmodel.Admin, error) {
	selectedAdmins, err := a.Queries.GetAllAdmins(*c)

	if err != nil {
		return nil, err
	}
	return &selectedAdmins, nil
}

func (a *PostgresAdminRepository) Save(c *context.Context, id *pgtype.UUID) (*dbmodel.CreateAdminRow, error) {
	selectedAdmin, err := a.Queries.CreateAdmin(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedAdmin, nil
}

func (a *PostgresAdminRepository) GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.Admin, error) {
	selectedAdmin, err := a.Queries.GetAdminByID(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedAdmin, nil
}

func (a *PostgresAdminRepository) Update(c *context.Context, id *pgtype.UUID) (*dbmodel.UpdateAdminRow, error) {
	selectedAdmin, err := a.Queries.UpdateAdmin(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedAdmin, nil
}

func (a *PostgresAdminRepository) Delete(c *context.Context, id *pgtype.UUID) (error) {
	if err := a.Queries.DeleteAdmin(*c, *id); err != nil {
		return err
	}
	return nil
}