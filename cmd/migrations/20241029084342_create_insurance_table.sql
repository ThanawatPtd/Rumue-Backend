-- +goose Up
-- +goose StatementBegin
CREATE TABLE "insurance" (
    brand VARCHAR(100),
    model VARCHAR(100),
    year VARCHAR(100),
    price FLOAT NOT NULL,

    PRIMARY KEY(brand, model, year)
);

INSERT INTO "insurance" (brand, model, year, price) VALUES
    ('Toyota', 'Camry', '2022', 1200),
    ('Toyota', 'Camry', '2021', 1150),
    ('Toyota', 'Corolla', '2023', 1050),
    ('Toyota', 'Corolla', '2022', 1000),
    ('Toyota', 'RAV4', '2023', 1300),
    ('Toyota', 'RAV4', '2021', 1250),

    ('Honda', 'Civic', '2023', 1100),
    ('Honda', 'Civic', '2021', 1000),
    ('Honda', 'Accord', '2022', 1250),
    ('Honda', 'Accord', '2021', 1200),
    ('Honda', 'CR-V', '2023', 1350),
    ('Honda', 'CR-V', '2022', 1300),

    ('Ford', 'Mustang', '2023', 1800),
    ('Ford', 'Mustang', '2022', 1750),
    ('Ford', 'Explorer', '2023', 1600),
    ('Ford', 'Explorer', '2021', 1550),
    ('Ford', 'F-150', '2023', 1400),
    ('Ford', 'F-150', '2022', 1350),

    ('BMW', '3 Series', '2023', 2600),
    ('BMW', '3 Series', '2022', 2500),
    ('BMW', 'X5', '2023', 3000),
    ('BMW', 'X5', '2021', 2900),
    ('BMW', '5 Series', '2023', 3200),
    ('BMW', '5 Series', '2022', 3100),

    ('Mercedes-Benz', 'C-Class', '2023', 3100),
    ('Mercedes-Benz', 'C-Class', '2022', 3000),
    ('Mercedes-Benz', 'E-Class', '2023', 3500),
    ('Mercedes-Benz', 'E-Class', '2021', 3400),
    ('Mercedes-Benz', 'GLC', '2023', 3300),
    ('Mercedes-Benz', 'GLC', '2022', 3200),

    ('Nissan', 'Altima', '2023', 1050),
    ('Nissan', 'Altima', '2021', 1000),
    ('Nissan', 'Maxima', '2022', 1150),
    ('Nissan', 'Maxima', '2021', 1100),
    ('Nissan', 'Rogue', '2023', 1250),
    ('Nissan', 'Rogue', '2022', 1200),

    ('Audi', 'A4', '2023', 2400),
    ('Audi', 'A4', '2021', 2300),
    ('Audi', 'Q5', '2023', 2800),
    ('Audi', 'Q5', '2022', 2700),
    ('Audi', 'A6', '2023', 3000),
    ('Audi', 'A6', '2021', 2900),

    ('Hyundai', 'Elantra', '2023', 950),
    ('Hyundai', 'Elantra', '2022', 900),
    ('Hyundai', 'Sonata', '2023', 1000),
    ('Hyundai', 'Sonata', '2022', 950),
    ('Hyundai', 'Tucson', '2023', 1100),
    ('Hyundai', 'Tucson', '2021', 1050);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "insurance";
-- +goose StatementEnd
