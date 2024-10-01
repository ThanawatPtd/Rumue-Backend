package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresInvoiceRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresInvoiceRepository(db *pgxpool.Pool) repositories.InvoiceRepository {
	return &PostgresInvoiceRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

func (p *PostgresInvoiceRepository) Save(c *context.Context, invoice *entities.Invoice) (*entities.Invoice, error) {
	var createInvoiceParams dbmodel.CreateInvoiceParams
	createInvoiceParams = dbmodel.CreateInvoiceParams{
		TransactionID:     convert.StringToUUID(invoice.TransactionID),
		Price:             invoice.Price,
		Invoice_image_url: invoice.Invoice_image_url,
	}
	createdInvoiceRow, err := p.Queries.CreateInvoice(*c, createInvoiceParams)

	if err != nil {
		return nil, err
	}

	var invoiceRes entities.Invoice
	invoiceRes = entities.Invoice{
		TransactionID:      convert.UUIDToString(createdInvoiceRow.TransactionID),
		Price:             createdInvoiceRow.Price,
		Invoice_image_url: createdInvoiceRow.Invoice_image_url,
	}

	return &invoiceRes, nil
}

func (i *PostgresInvoiceRepository) ListAll(c *context.Context) (*[]dbmodel.GetAllInvoicesRow, error) {
	selectedInvoices, err := i.Queries.GetAllInvoices(*c)

	if err != nil {
		return nil, err
	}

	return &selectedInvoices, nil
}
