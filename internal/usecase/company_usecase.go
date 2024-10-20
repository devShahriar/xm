package usecase

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/devShahriar/H"
	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/entity"
	"github.com/devShahriar/xm/internal/repository"
	"github.com/google/uuid"
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

func (uc *CompanyUsecase) CreateCompany(ctx context.Context, company *entity.Company) error {
	// Perform validations on the company data
	if err := ValidateCompany(company); err != nil {
		return err
	}

	// Call the repository to persist the company data
	return uc.repo.CreateCompany(ctx, company)
}

func (uc *CompanyUsecase) UpdateCompany(ctx context.Context, company *entity.Company) (*entity.Company, error) {
	// Check if the company exists
	existingCompany, err := uc.repo.GetCompanyByID(ctx, company.ID.String())
	if err != nil {
		if err == common.ErrCompanyNotFound {
			return nil, common.ErrCompanyNotFound
		}
		return nil, err
	}
	spew.Dump(existingCompany)
	// Update the fields (for simplicity, we overwrite all updatable fields)
	existingCompany.Name = H.If(company.Name != "", company.Name, existingCompany.Name)
	existingCompany.Description = H.If(company.Description != "", company.Description, existingCompany.Description)
	existingCompany.NumEmployees = H.If(company.NumEmployees != 0, company.NumEmployees, existingCompany.NumEmployees)
	existingCompany.Registered = company.Registered
	existingCompany.Type = H.If(company.Type != "", company.Type, existingCompany.Type)

	if err := ValidateCompany(existingCompany); err != nil {
		log.Printf("validate company error %v", err)
		return nil, err
	}
	spew.Dump(existingCompany)
	// Save the updated company in the repository
	return existingCompany, uc.repo.UpdateCompany(ctx, existingCompany)
}

func (uc *CompanyUsecase) DeleteCompany(ctx context.Context, companyID uuid.UUID) error {
	// Check if the company exists
	_, err := uc.repo.GetCompanyByID(ctx, companyID.String())
	if err != nil {
		if err == common.ErrCompanyNotFound {
			return common.ErrCompanyNotFound
		}
		return err
	}

	// Call the repository to delete the company
	return uc.repo.DeleteCompany(ctx, companyID.String())
}
