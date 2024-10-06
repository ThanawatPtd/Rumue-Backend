-- +goose Up
-- +goose StatementBegin
-- Invoice Table
CREATE TABLE "invoice" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL,
    price FLOAT NOT NULL,
    invoice_image_url VARCHAR(100) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES "transaction"(id) ON DELETE CASCADE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "invoice";
-- +goose StatementEnd
