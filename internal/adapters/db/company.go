package db

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"

	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/config"
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

func CreateDBIfNotExists() error {
	conf := config.GetAppConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		conf.DbConfig.Host,
		conf.DbConfig.User,
		conf.DbConfig.Password,
		conf.DbConfig.Port,
	)
	dbName := config.GetAppConfig().DbConfig.DbName

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to PostgreSQL server: %v", err)
	}

	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	fmt.Printf("Database '%s' is ready.\n", dbName)
	return nil
}

func (db *CompanyDB) CreateCompany(ctx context.Context, company *entity.Company) error {
	return db.Db.WithContext(ctx).Create(company).Error
}

func (db *CompanyDB) UpdateCompany(ctx context.Context, company *entity.Company) error {
	return db.Db.WithContext(ctx).Save(company).Error
}

func (db *CompanyDB) DeleteCompany(ctx context.Context, id string) error {
	return db.Db.WithContext(ctx).Delete(&entity.Company{}, "id = ?", id).Error
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
