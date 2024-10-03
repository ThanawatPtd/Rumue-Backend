-- name: GetAllInvoices :many
SELECT
    id,
    transaction_id,
    price,
    invoice_image_url
from "invoice";

-- name: GetInvoiceByID :one
SELECT
    id,
    transaction_id,
    price,
    invoice_image_url
from "invoice"
WHERE id = $1;

-- name: CreateInvoice :one
INSERT INTO "invoice" (
    transaction_id, price, invoice_image_url
) VALUES (
    $1, $2, $3
)
RETURNING id, transaction_id, price, invoice_image_url;
