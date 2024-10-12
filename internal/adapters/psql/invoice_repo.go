package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
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

func (p *PostgresInvoiceRepository) Save(c context.Context, invoice *entities.Invoice) (*entities.Invoice, error) {
	var createInvoiceParams dbmodel.CreateInvoiceParams
	createInvoiceParams = dbmodel.CreateInvoiceParams{
		TransactionID:     convert.StringToUUID(invoice.TransactionID),
		Price:             invoice.Price,
		InvoiceImageUrl: invoice.Invoice_image_url,
	}
	createdInvoiceRow, err := p.Queries.CreateInvoice(c, createInvoiceParams)

	if err != nil {
		return nil, err
	}

	var invoiceRes entities.Invoice
	invoiceRes = entities.Invoice{
		TransactionID:      convert.UUIDToString(createdInvoiceRow.TransactionID),
		Price:             createdInvoiceRow.Price,
		Invoice_image_url: createdInvoiceRow.InvoiceImageUrl,
	}

	return &invoiceRes, nil
}

func (i *PostgresInvoiceRepository) ListAll(c context.Context) ([]entities.Invoice, error) {
	selectedInvoices, err := i.Queries.GetAllInvoices(c)
	if err != nil {
		return nil, err
	}

	invoices := []entities.Invoice{}
	for _, value := range selectedInvoices {
		invoice := entities.Invoice{}
		if err := utils.MappingParser(&value, &invoice); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
