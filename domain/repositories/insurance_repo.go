package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type InsuranceRepository interface{
	GetInsurances(c context.Context) ([]entities.Insurance, error)
	GetInsurance(c context.Context, rep *entities.Insurance) (*entities.Insurance, error)
}