package usecase

import (
	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/entity"
)

func ValidateCompany(company *entity.Company) error {
	// Validate company name (required, max 15 characters)
	if len(company.Name) == 0 || len(company.Name) > 15 {
		return common.ErrInvalidCompanyName
	}

	// Validate description (optional, max 3000 characters)
	if len(company.Description) > 3000 {
		return common.ErrInvalidDescription
	}

	// Validate number of employees (required, must be at least 1)
	if company.NumEmployees < 1 {
		return common.ErrInvalidNumEmployees
	}

	// Validate company type (required, must be one of the predefined types)
	if company.Type == "" {
		return common.ErrCompanyTypeRequired
	}

	if !ValidateCompanyType(company.Type) {
		return common.ErrInvalidCompanyType
	}

	return nil
}

func ValidateCompanyType(companyType string) bool {
	validCompanyTypes := common.GetValidCompanyTypes()
	return validCompanyTypes[companyType]
}
