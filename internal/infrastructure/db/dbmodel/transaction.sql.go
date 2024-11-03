// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transaction.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO "transaction" (
    user_id, vehicle_id, price, insurance_type, status, e_slip_image_url, cr_image_url, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
)
RETURNING id, user_id, vehicle_id, price, insurance_type, status, e_slip_image_url, cr_image_url, cip_number,
vip_number
`

type CreateTransactionParams struct {
	UserID        pgtype.UUID `json:"userId"`
	VehicleID     pgtype.UUID `json:"vehicleId"`
	Price         float64     `json:"price"`
	InsuranceType string      `json:"insuranceType"`
	Status        string      `json:"status"`
	ESlipImageUrl string      `json:"eSlipImageUrl"`
	CrImageUrl    string      `json:"crImageUrl"`
}

type CreateTransactionRow struct {
	ID            pgtype.UUID `json:"id"`
	UserID        pgtype.UUID `json:"userId"`
	VehicleID     pgtype.UUID `json:"vehicleId"`
	Price         float64     `json:"price"`
	InsuranceType string      `json:"insuranceType"`
	Status        string      `json:"status"`
	ESlipImageUrl string      `json:"eSlipImageUrl"`
	CrImageUrl    string      `json:"crImageUrl"`
	CipNumber     pgtype.Text `json:"cipNumber"`
	VipNumber     pgtype.Text `json:"vipNumber"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (CreateTransactionRow, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.UserID,
		arg.VehicleID,
		arg.Price,
		arg.InsuranceType,
		arg.Status,
		arg.ESlipImageUrl,
		arg.CrImageUrl,
	)
	var i CreateTransactionRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VehicleID,
		&i.Price,
		&i.InsuranceType,
		&i.Status,
		&i.ESlipImageUrl,
		&i.CrImageUrl,
		&i.CipNumber,
		&i.VipNumber,
	)
	return i, err
}

const getAllTransactions = `-- name: GetAllTransactions :many
SELECT
    id,
    user_id,
    vehicle_id,
    employee_id,
    price,
    insurance_type,
    status,
    e_slip_image_url,
    cr_image_url,
    created_at,
    updated_at 
FROM "transaction"
`

type GetAllTransactionsRow struct {
	ID            pgtype.UUID        `json:"id"`
	UserID        pgtype.UUID        `json:"userId"`
	VehicleID     pgtype.UUID        `json:"vehicleId"`
	EmployeeID    pgtype.UUID        `json:"employeeId"`
	Price         float64            `json:"price"`
	InsuranceType string             `json:"insuranceType"`
	Status        string             `json:"status"`
	ESlipImageUrl string             `json:"eSlipImageUrl"`
	CrImageUrl    string             `json:"crImageUrl"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt     pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetAllTransactions(ctx context.Context) ([]GetAllTransactionsRow, error) {
	rows, err := q.db.Query(ctx, getAllTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTransactionsRow
	for rows.Next() {
		var i GetAllTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VehicleID,
			&i.EmployeeID,
			&i.Price,
			&i.InsuranceType,
			&i.Status,
			&i.ESlipImageUrl,
			&i.CrImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getAllTransactionsByUserID = `-- name: GetAllTransactionsByUserID :many
SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.user_id = $1
`

type GetAllTransactionsByUserIDRow struct {
	ID                    pgtype.UUID        `json:"id"`
	UserID                pgtype.UUID        `json:"userId"`
	VehicleID             pgtype.UUID        `json:"vehicleId"`
	ReceiptDate           pgtype.Timestamptz `json:"receiptDate"`
	Email                 string             `json:"email"`
	Fname                 string             `json:"fname"`
	Lname                 string             `json:"lname"`
	PhoneNumber           string             `json:"phoneNumber"`
	Address               string             `json:"address"`
	Nationality           string             `json:"nationality"`
	BirthDate             pgtype.Timestamptz `json:"birthDate"`
	CitizenID             string             `json:"citizenId"`
	RegistrationDate      pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber    string             `json:"registrationNumber"`
	Province              string             `json:"province"`
	VehicleType           string             `json:"vehicleType"`
	VehicleCategory       string             `json:"vehicleCategory"`
	Characteristics       string             `json:"characteristics"`
	Brand                 string             `json:"brand"`
	Model                 string             `json:"model"`
	ModelYear             string             `json:"modelYear"`
	VehicleColor          string             `json:"vehicleColor"`
	VehicleNumber         string             `json:"vehicleNumber"`
	VehicleNumberLocation string             `json:"vehicleNumberLocation"`
	EngineBrand           string             `json:"engineBrand"`
	EngineNumber          string             `json:"engineNumber"`
	EngineNumberLocation  string             `json:"engineNumberLocation"`
	FuelType              string             `json:"fuelType"`
	ChasisNumber          string             `json:"chasisNumber"`
	WheelType             string             `json:"wheelType"`
	TotalPiston           int32              `json:"totalPiston"`
	Cc                    int32              `json:"cc"`
	HorsePower            int32              `json:"horsePower"`
	WeightUnlanden        float64            `json:"weightUnlanden"`
	WeightLaden           float64            `json:"weightLaden"`
	SeatingCapacity       int32              `json:"seatingCapacity"`
	Miles                 float64            `json:"miles"`
	InsuranceType         string             `json:"insuranceType"`
	Status                string             `json:"status"`
	ESlipImageUrl         string             `json:"eSlipImageUrl"`
	CrImageUrl            string             `json:"crImageUrl"`
	CipNumber             pgtype.Text        `json:"cipNumber"`
	VipNumber             pgtype.Text        `json:"vipNumber"`
	Price                 float64            `json:"price"`
	CreatedAt             pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt             pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetAllTransactionsByUserID(ctx context.Context, userID pgtype.UUID) ([]GetAllTransactionsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, getAllTransactionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTransactionsByUserIDRow
	for rows.Next() {
		var i GetAllTransactionsByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VehicleID,
			&i.ReceiptDate,
			&i.Email,
			&i.Fname,
			&i.Lname,
			&i.PhoneNumber,
			&i.Address,
			&i.Nationality,
			&i.BirthDate,
			&i.CitizenID,
			&i.RegistrationDate,
			&i.RegistrationNumber,
			&i.Province,
			&i.VehicleType,
			&i.VehicleCategory,
			&i.Characteristics,
			&i.Brand,
			&i.Model,
			&i.ModelYear,
			&i.VehicleColor,
			&i.VehicleNumber,
			&i.VehicleNumberLocation,
			&i.EngineBrand,
			&i.EngineNumber,
			&i.EngineNumberLocation,
			&i.FuelType,
			&i.ChasisNumber,
			&i.WheelType,
			&i.TotalPiston,
			&i.Cc,
			&i.HorsePower,
			&i.WeightUnlanden,
			&i.WeightLaden,
			&i.SeatingCapacity,
			&i.Miles,
			&i.InsuranceType,
			&i.Status,
			&i.ESlipImageUrl,
			&i.CrImageUrl,
			&i.CipNumber,
			&i.VipNumber,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getExpiredInsuranceTransactions = `-- name: GetExpiredInsuranceTransactions :many

SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
    u.email, u.fname, u.lname, u.phone_number,
    u.address, u.nationality, u.birth_date, u.citizen_id,
    v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, 
    v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, 
    v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, 
    v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, 
    v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, 
    t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" AS t
JOIN "vehicle" AS v ON t.vehicle_id = v.id 
JOIN "user" AS u ON t.user_id = u.id
WHERE t.updated_at + INTERVAL '1 year' 
      BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '1 week'
`

type GetExpiredInsuranceTransactionsRow struct {
	ID                    pgtype.UUID        `json:"id"`
	UserID                pgtype.UUID        `json:"userId"`
	VehicleID             pgtype.UUID        `json:"vehicleId"`
	ReceiptDate           pgtype.Timestamptz `json:"receiptDate"`
	Email                 string             `json:"email"`
	Fname                 string             `json:"fname"`
	Lname                 string             `json:"lname"`
	PhoneNumber           string             `json:"phoneNumber"`
	Address               string             `json:"address"`
	Nationality           string             `json:"nationality"`
	BirthDate             pgtype.Timestamptz `json:"birthDate"`
	CitizenID             string             `json:"citizenId"`
	RegistrationDate      pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber    string             `json:"registrationNumber"`
	Province              string             `json:"province"`
	VehicleType           string             `json:"vehicleType"`
	VehicleCategory       string             `json:"vehicleCategory"`
	Characteristics       string             `json:"characteristics"`
	Brand                 string             `json:"brand"`
	Model                 string             `json:"model"`
	ModelYear             string             `json:"modelYear"`
	VehicleColor          string             `json:"vehicleColor"`
	VehicleNumber         string             `json:"vehicleNumber"`
	VehicleNumberLocation string             `json:"vehicleNumberLocation"`
	EngineBrand           string             `json:"engineBrand"`
	EngineNumber          string             `json:"engineNumber"`
	EngineNumberLocation  string             `json:"engineNumberLocation"`
	FuelType              string             `json:"fuelType"`
	ChasisNumber          string             `json:"chasisNumber"`
	WheelType             string             `json:"wheelType"`
	TotalPiston           int32              `json:"totalPiston"`
	Cc                    int32              `json:"cc"`
	HorsePower            int32              `json:"horsePower"`
	WeightUnlanden        float64            `json:"weightUnlanden"`
	WeightLaden           float64            `json:"weightLaden"`
	SeatingCapacity       int32              `json:"seatingCapacity"`
	Miles                 float64            `json:"miles"`
	InsuranceType         string             `json:"insuranceType"`
	Status                string             `json:"status"`
	ESlipImageUrl         string             `json:"eSlipImageUrl"`
	CrImageUrl            string             `json:"crImageUrl"`
	CipNumber             pgtype.Text        `json:"cipNumber"`
	VipNumber             pgtype.Text        `json:"vipNumber"`
	Price                 float64            `json:"price"`
	CreatedAt             pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt             pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetExpiredInsuranceTransactions(ctx context.Context) ([]GetExpiredInsuranceTransactionsRow, error) {
	rows, err := q.db.Query(ctx, getExpiredInsuranceTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetExpiredInsuranceTransactionsRow
	for rows.Next() {
		var i GetExpiredInsuranceTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VehicleID,
			&i.ReceiptDate,
			&i.Email,
			&i.Fname,
			&i.Lname,
			&i.PhoneNumber,
			&i.Address,
			&i.Nationality,
			&i.BirthDate,
			&i.CitizenID,
			&i.RegistrationDate,
			&i.RegistrationNumber,
			&i.Province,
			&i.VehicleType,
			&i.VehicleCategory,
			&i.Characteristics,
			&i.Brand,
			&i.Model,
			&i.ModelYear,
			&i.VehicleColor,
			&i.VehicleNumber,
			&i.VehicleNumberLocation,
			&i.EngineBrand,
			&i.EngineNumber,
			&i.EngineNumberLocation,
			&i.FuelType,
			&i.ChasisNumber,
			&i.WheelType,
			&i.TotalPiston,
			&i.Cc,
			&i.HorsePower,
			&i.WeightUnlanden,
			&i.WeightLaden,
			&i.SeatingCapacity,
			&i.Miles,
			&i.InsuranceType,
			&i.Status,
			&i.ESlipImageUrl,
			&i.CrImageUrl,
			&i.CipNumber,
			&i.VipNumber,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getTransactionByID = `-- name: GetTransactionByID :one
SELECT 
    id,
    user_id,
    vehicle_id,
    employee_id,
    price,
    insurance_type,
    status,
    e_slip_image_url,
    cr_image_url,
    receipt_date,
    created_at,
    updated_at 
FROM "transaction"
WHERE id = $1
`

type GetTransactionByIDRow struct {
	ID            pgtype.UUID        `json:"id"`
	UserID        pgtype.UUID        `json:"userId"`
	VehicleID     pgtype.UUID        `json:"vehicleId"`
	EmployeeID    pgtype.UUID        `json:"employeeId"`
	Price         float64            `json:"price"`
	InsuranceType string             `json:"insuranceType"`
	Status        string             `json:"status"`
	ESlipImageUrl string             `json:"eSlipImageUrl"`
	CrImageUrl    string             `json:"crImageUrl"`
	ReceiptDate   pgtype.Timestamptz `json:"receiptDate"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt     pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetTransactionByID(ctx context.Context, id pgtype.UUID) (GetTransactionByIDRow, error) {
	row := q.db.QueryRow(ctx, getTransactionByID, id)
	var i GetTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VehicleID,
		&i.EmployeeID,
		&i.Price,
		&i.InsuranceType,
		&i.Status,
		&i.ESlipImageUrl,
		&i.CrImageUrl,
		&i.ReceiptDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserVehicleTransactionByID = `-- name: GetUserVehicleTransactionByID :one
SELECT
    t.id, t.user_id, t.vehicle_id, t.receipt_date,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.id = $1
`

type GetUserVehicleTransactionByIDRow struct {
	ID                    pgtype.UUID        `json:"id"`
	UserID                pgtype.UUID        `json:"userId"`
	VehicleID             pgtype.UUID        `json:"vehicleId"`
	ReceiptDate           pgtype.Timestamptz `json:"receiptDate"`
	Email                 string             `json:"email"`
	Fname                 string             `json:"fname"`
	Lname                 string             `json:"lname"`
	PhoneNumber           string             `json:"phoneNumber"`
	Address               string             `json:"address"`
	Nationality           string             `json:"nationality"`
	BirthDate             pgtype.Timestamptz `json:"birthDate"`
	CitizenID             string             `json:"citizenId"`
	RegistrationDate      pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber    string             `json:"registrationNumber"`
	Province              string             `json:"province"`
	VehicleType           string             `json:"vehicleType"`
	VehicleCategory       string             `json:"vehicleCategory"`
	Characteristics       string             `json:"characteristics"`
	Brand                 string             `json:"brand"`
	Model                 string             `json:"model"`
	ModelYear             string             `json:"modelYear"`
	VehicleColor          string             `json:"vehicleColor"`
	VehicleNumber         string             `json:"vehicleNumber"`
	VehicleNumberLocation string             `json:"vehicleNumberLocation"`
	EngineBrand           string             `json:"engineBrand"`
	EngineNumber          string             `json:"engineNumber"`
	EngineNumberLocation  string             `json:"engineNumberLocation"`
	FuelType              string             `json:"fuelType"`
	ChasisNumber          string             `json:"chasisNumber"`
	WheelType             string             `json:"wheelType"`
	TotalPiston           int32              `json:"totalPiston"`
	Cc                    int32              `json:"cc"`
	HorsePower            int32              `json:"horsePower"`
	WeightUnlanden        float64            `json:"weightUnlanden"`
	WeightLaden           float64            `json:"weightLaden"`
	SeatingCapacity       int32              `json:"seatingCapacity"`
	Miles                 float64            `json:"miles"`
	InsuranceType         string             `json:"insuranceType"`
	Status                string             `json:"status"`
	ESlipImageUrl         string             `json:"eSlipImageUrl"`
	CrImageUrl            string             `json:"crImageUrl"`
	CipNumber             pgtype.Text        `json:"cipNumber"`
	VipNumber             pgtype.Text        `json:"vipNumber"`
	Price                 float64            `json:"price"`
	CreatedAt             pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt             pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetUserVehicleTransactionByID(ctx context.Context, id pgtype.UUID) (GetUserVehicleTransactionByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserVehicleTransactionByID, id)
	var i GetUserVehicleTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VehicleID,
		&i.ReceiptDate,
		&i.Email,
		&i.Fname,
		&i.Lname,
		&i.PhoneNumber,
		&i.Address,
		&i.Nationality,
		&i.BirthDate,
		&i.CitizenID,
		&i.RegistrationDate,
		&i.RegistrationNumber,
		&i.Province,
		&i.VehicleType,
		&i.VehicleCategory,
		&i.Characteristics,
		&i.Brand,
		&i.Model,
		&i.ModelYear,
		&i.VehicleColor,
		&i.VehicleNumber,
		&i.VehicleNumberLocation,
		&i.EngineBrand,
		&i.EngineNumber,
		&i.EngineNumberLocation,
		&i.FuelType,
		&i.ChasisNumber,
		&i.WheelType,
		&i.TotalPiston,
		&i.Cc,
		&i.HorsePower,
		&i.WeightUnlanden,
		&i.WeightLaden,
		&i.SeatingCapacity,
		&i.Miles,
		&i.InsuranceType,
		&i.Status,
		&i.ESlipImageUrl,
		&i.CrImageUrl,
		&i.CipNumber,
		&i.VipNumber,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateReceiptDateTransacton = `-- name: UpdateReceiptDateTransacton :one
UPDATE "transaction"
SET
    receipt_date = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING receipt_date
`

func (q *Queries) UpdateReceiptDateTransacton(ctx context.Context, id pgtype.UUID) (pgtype.Timestamptz, error) {
	row := q.db.QueryRow(ctx, updateReceiptDateTransacton, id)
	var receipt_date pgtype.Timestamptz
	err := row.Scan(&receipt_date)
	return receipt_date, err
}

const updateTransaction = `-- name: UpdateTransaction :exec
UPDATE "transaction"
SET
    employee_id = $2,
    status = $3,
    cip_number = $4,
    vip_number = $5,
    updated_at = NOW()
WHERE id = $1
`

type UpdateTransactionParams struct {
	ID         pgtype.UUID `json:"id"`
	EmployeeID pgtype.UUID `json:"employeeId"`
	Status     string      `json:"status"`
	CipNumber  pgtype.Text `json:"cipNumber"`
	VipNumber  pgtype.Text `json:"vipNumber"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.db.Exec(ctx, updateTransaction,
		arg.ID,
		arg.EmployeeID,
		arg.Status,
		arg.CipNumber,
		arg.VipNumber,
	)
	return err
}
