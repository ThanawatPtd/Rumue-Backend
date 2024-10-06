package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	// "github.com/jackc/pgx/v5/pgtype"
)

type InvoiceUseCase interface {
	CreateInvoice(ctx context.Context, invoice *entities.Invoice) (*entities.Invoice, error)
	// GetInvoiceByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetInvoiceByIDRow, error)
	// DeleteByID(ctx context.Context, id *pgtype.UUID) error
	GetInvoices(ctx context.Context) (*[]entities.Invoice, error)
	// UpdateUser(c context.Context, id *pgtype.UUID, user *dbmodel.UpdateUserParams) (*dbmodel.UpdateUserRow, error)
}

type InvoiceService struct {
	invoiceRepo repositories.InvoiceRepository
}

func ProvideInvoiceService(invoiceRepo repositories.InvoiceRepository) InvoiceUseCase {
	return &InvoiceService{
		invoiceRepo: invoiceRepo,
	}
}

// func (i *InvoiceService) GetInvoiceByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.GetInvoiceByIDRow, error) {
// 	selected, err := i.invoiceRepo.GetByID(&ctx, id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return selected, nil
// }

func (i *InvoiceService) GetInvoices(ctx context.Context) (*[]entities.Invoice, error) {
	list, err := i.invoiceRepo.ListAll(&ctx)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (i *InvoiceService) CreateInvoice(ctx context.Context, invoice *entities.Invoice) (*entities.Invoice, error) {

	newInvoice, err := i.invoiceRepo.Save(&ctx, invoice)

	if err != nil {
		return nil, err
	}

	return newInvoice, nil
}

// func (i *InvoiceService) DeleteByID(ctx context.Context, id *pgtype.UUID) error {
// 	if err := i.invoiceRepo.Delete(&ctx, id); err != nil {
// 		return err
// 	}

// 	return nil
// }
