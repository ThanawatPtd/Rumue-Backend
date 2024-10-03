-- name: GetAllInvoices :many
SELECT
    invoice_id,
    transaction_id,
    price,
    invoice_image_url
from "invoice";

-- name: GetInvoiceByID :one
SELECT
    invoice_id,
    transaction_id,
    price,
    invoice_image_url
from "invoice"
WHERE invoice_id = $1;

-- name: CreateInvoice :one
INSERT INTO "invoice" (
    transaction_id, price, invoice_image_url
) VALUES (
    $1, $2, $3
)
RETURNING invoice_id, transaction_id, price, invoice_image_url;
