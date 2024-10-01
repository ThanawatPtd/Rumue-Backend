package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	// "github.com/jackc/pgx/v5/pgtype"
)

type InvoiceRepository interface {
	ListAll(c *context.Context) (*[]dbmodel.GetAllInvoicesRow, error)
	Save(c *context.Context, invoice *entities.Invoice) (*entities.Invoice, error)
	// GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.GetInvoiceByIDRow, error)
	// Update(c *context.Context, user *dbmodel.UpdateInvoiceParams) (*dbmodel.UpdateInvoiceRow, error)
	// Delete(c *context.Context, id *pgtype.UUID) error
}
