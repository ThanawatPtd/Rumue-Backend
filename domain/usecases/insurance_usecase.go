package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/domain/responses"
)

type InsuranceUseCase interface {
	GetInsurance(c context.Context, req *entities.Insurance) (*entities.Insurance, error)
	GetInsurances(c context.Context) (*responses.GetInsurancesResponse, error)
}

type InsuranceService struct {
	repo repositories.InsuranceRepository
}

func ProvideInsuranceService(repo repositories.InsuranceRepository) InsuranceUseCase {
	return &InsuranceService{
		repo: repo,
	}
}

func (i *InsuranceService) GetInsurance(c context.Context, req *entities.Insurance) (*entities.Insurance, error) {
	return i.repo.GetInsurance(c, req)
}

func (i *InsuranceService) GetInsurances(c context.Context) (*responses.GetInsurancesResponse, error) {
	list, err := i.repo.GetInsurances(c)
	if err != nil{
		return nil, err
	}

	tree := make(map[string]map[string][]string)

	for _, val := range list {
		if _, exist := tree[val.Brand]; !exist{
			tree[val.Brand] = make(map[string][]string)
		}

		if _, exist := tree[val.Brand][val.Model]; !exist{
			tree[val.Brand][val.Model] = []string{}
		}

		tree[val.Brand][val.Model] = append(tree[val.Brand][val.Model], val.Year)
	}

	response := &responses.GetInsurancesResponse{
		Tree: tree,
	}

	return response, nil

}


