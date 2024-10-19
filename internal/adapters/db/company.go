package db

import (
	"context"

	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/entity"
	"github.com/devShahriar/xm/internal/repository"
	"gorm.io/gorm"
)

var _ repository.CompanyRepository = &CompanyDB{} // This will ensure that CompanyDB has successfully implemented CompanyRepository

type CompanyDB struct {
	Db *gorm.DB
}

func (d *CompanyDB) RunMigrations() {
	d.Db.AutoMigrate(
		entity.Company{},
	)
}

func (db *CompanyDB) CreateCompany(ctx context.Context, company *entity.Company) error {
	return nil
}

func (db *CompanyDB) UpdateCompany(ctx context.Context, company *entity.Company) error {
	return nil
}

func (db *CompanyDB) DeleteCompany(ctx context.Context, id string) error {
	return nil
}

func (db *CompanyDB) GetCompanyByID(ctx context.Context, id string) (*entity.Company, error) {
	var company entity.Company
	err := db.Db.WithContext(ctx).First(&company, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrCompanyNotFound // Handle the case when the company is not found
		}
		return nil, err
	}
	return &company, nil
}

func (db *CompanyDB) GetCompanies(ctx context.Context) ([]*entity.Company, error) {
	return nil, nil
}
