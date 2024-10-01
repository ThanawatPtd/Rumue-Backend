package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllInvoices = `-- name: GetAllInvoices :many
SELECT
    invoice_id,
    transaction_id,
    price,
    invoice_image_url
from "invoice";
`

type GetAllInvoicesRow struct {
	InvoiceID         pgtype.UUID `json:"invoiceId"`
	TransactionID     pgtype.UUID `json:"transactionId"`
	Price             float64     `json:"price"`
	Invoice_image_url string      `json:"invoiceImageUrl"`
}

func (q *Queries) GetAllInvoices(ctx context.Context) ([]GetAllInvoicesRow, error) {
	rows, err := q.db.Query(ctx, getAllInvoices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllInvoicesRow
	for rows.Next() {
		var i GetAllInvoicesRow
		if err := rows.Scan(
			&i.InvoiceID,
			&i.TransactionID,
			&i.Price,
			&i.Invoice_image_url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO "invoice" (
    transaction_id, price, invoice_image_url
) VALUES (
    $1, $2, $3
)
RETURNING invoice_id, transaction_id, price, invoice_image_url;

`

type CreateInvoiceParams struct {
	TransactionID     pgtype.UUID `json:"transactionId"`
	Price             float64     `json:"price"`
	Invoice_image_url string      `json:"invoiceImageUrl"`
}

type CreateInvoiceRow struct {
	InvoiceID         pgtype.UUID `json:"invoiceId"`
	TransactionID     pgtype.UUID `json:"transactionId"`
	Price             float64     `json:"price"`
	Invoice_image_url string      `json:"invoiceImageUrl"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (CreateInvoiceRow, error) {
	row := q.db.QueryRow(ctx, createInvoice,
		arg.TransactionID,
		arg.Price,
		arg.Invoice_image_url,
	)

	var i CreateInvoiceRow
	err := row.Scan(
		&i.InvoiceID,
		&i.TransactionID,
		&i.Price,
		&i.Invoice_image_url,
	)

	return i, err
}
