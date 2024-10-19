package repository

import (
	"context"

	"github.com/devShahriar/xm/internal/entity"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	UpdateCompany(ctx context.Context, company *entity.Company) error
	DeleteCompany(ctx context.Context, id string) error
	GetCompanyByID(ctx context.Context, id string) (*entity.Company, error)
	GetCompanies(ctx context.Context) ([]*entity.Company, error)
}
