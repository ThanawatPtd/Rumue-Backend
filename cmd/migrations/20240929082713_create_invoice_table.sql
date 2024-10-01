-- +goose Up
-- +goose StatementBegin
-- Invoice Table
CREATE TABLE "invoice" (
    invoice_id UUID PRIMARY KEY,
    transaction_id UUID NOT NULL,
    price FLOAT NOT NULL,
    invoice_image_url VARCHAR(100) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES "transaction"(transaction_id) ON DELETE CASCADE
);
-- +goose StatementEnd

ALTER TABLE "invoice" ALTER COLUMN invoice_id SET DEFAULT uuid_generate_v4();

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
