// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employee.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO "employee" (
    id, salary, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING id, salary, created_at
`

type CreateEmployeeParams struct {
	ID     pgtype.UUID `json:"id"`
	Salary float64     `json:"salary"`
}

type CreateEmployeeRow struct {
	ID        pgtype.UUID        `json:"id"`
	Salary    float64            `json:"salary"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (CreateEmployeeRow, error) {
	row := q.db.QueryRow(ctx, createEmployee, arg.ID, arg.Salary)
	var i CreateEmployeeRow
	err := row.Scan(&i.ID, &i.Salary, &i.CreatedAt)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM "employee"
WHERE id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteEmployee, id)
	return err
}

const findInsuranceToday = `-- name: FindInsuranceToday :many
SELECT
    t.id, t.user_id, t.vehicle_id,
u.email, u.fname, u.lname, u.phone_number,
u.address, u.nationality, u.birth_date, u.citizen_id,
v.registration_date, v.registration_number, v.province, v.vehicle_type, v.vehicle_category, v.characteristics, v.brand, v.model, v.model_year, v.vehicle_color, v.vehicle_number, v.vehicle_number_location, v.engine_brand, v.engine_number, v.engine_number_location, v.fuel_type, v.chasis_number, v.wheel_type, v.total_piston, v.cc, v.horse_power, v.weight_unlanden, v.weight_laden, v.seating_capacity, v.miles, t.insurance_type, t.status, t.e_slip_image_url, t.cr_image_url, t.cip_number, t.vip_number, t.price, t.created_at, t.updated_at
FROM "transaction" as t
JOIN "vehicle" as v ON t.vehicle_id = v.id 
JOIN "user" as u ON t.user_id = u.id
WHERE t.status = 'pending'
or t.updated_at::date = CURRENT_DATE
`

type FindInsuranceTodayRow struct {
	ID                    pgtype.UUID        `json:"id"`
	UserID                pgtype.UUID        `json:"userId"`
	VehicleID             pgtype.UUID        `json:"vehicleId"`
	Email                 string             `json:"email"`
	Fname                 string             `json:"fname"`
	Lname                 string             `json:"lname"`
	PhoneNumber           string             `json:"phoneNumber"`
	Address               string             `json:"address"`
	Nationality           string             `json:"nationality"`
	BirthDate             pgtype.Date        `json:"birthDate"`
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

func (q *Queries) FindInsuranceToday(ctx context.Context) ([]FindInsuranceTodayRow, error) {
	rows, err := q.db.Query(ctx, findInsuranceToday)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindInsuranceTodayRow
	for rows.Next() {
		var i FindInsuranceTodayRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VehicleID,
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

const getAllEmployees = `-- name: GetAllEmployees :many
SELECT
   id,
   salary,
   created_at,
   updated_at 
FROM "employee"
`

func (q *Queries) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.Query(ctx, getAllEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.Salary,
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

const getEmployeeByID = `-- name: GetEmployeeByID :one
SELECT
    id,
    salary
FROM "employee"
WHERE id = $1
`

type GetEmployeeByIDRow struct {
	ID     pgtype.UUID `json:"id"`
	Salary float64     `json:"salary"`
}

func (q *Queries) GetEmployeeByID(ctx context.Context, id pgtype.UUID) (GetEmployeeByIDRow, error) {
	row := q.db.QueryRow(ctx, getEmployeeByID, id)
	var i GetEmployeeByIDRow
	err := row.Scan(&i.ID, &i.Salary)
	return i, err
}

const updateEmployee = `-- name: UpdateEmployee :one
UPDATE "employee"
SET
    salary = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, salary, updated_at
`

type UpdateEmployeeParams struct {
	ID     pgtype.UUID `json:"id"`
	Salary float64     `json:"salary"`
}

type UpdateEmployeeRow struct {
	ID        pgtype.UUID        `json:"id"`
	Salary    float64            `json:"salary"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (UpdateEmployeeRow, error) {
	row := q.db.QueryRow(ctx, updateEmployee, arg.ID, arg.Salary)
	var i UpdateEmployeeRow
	err := row.Scan(&i.ID, &i.Salary, &i.UpdatedAt)
	return i, err
}
