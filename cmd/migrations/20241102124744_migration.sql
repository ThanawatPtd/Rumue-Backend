-- +goose Up
-- +goose StatementBegin
ALTER TABLE "vehicle"
ADD CONSTRAINT fk_vehicle_brand_model_year 
FOREIGN KEY (brand, model, model_year) 
REFERENCES "insurance"(brand, model, year) 
ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
