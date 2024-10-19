package usecase

import (
	"context"

	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/entity"
	"github.com/devShahriar/xm/internal/repository"
)

type CompanyUsecase struct {
	repo repository.CompanyRepository
}

func NewCompanyUsecase(repo repository.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{repo: repo}
}

func (uc *CompanyUsecase) GetCompanyByID(ctx context.Context, id string) (*entity.Company, error) {

	// Fetch the company from the repository
	company, err := uc.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, common.ErrCompanyNotFound
	}

	return company, nil
}
