package common

import "errors"

var (
	ErrCompanyNotFound = errors.New("company not found")

	ErrInvalidCompanyName       = errors.New("company name is required and must be less than 15 characters")
	ErrInvalidDescription       = errors.New("company description must be less than 3000 characters")
	ErrInvalidNumEmployees      = errors.New("company must have at least one employee")
	ErrCompanyTypeRequired      = errors.New("company type is required")
	ErrInvalidCompanyType       = errors.New("invalid company type")
	ErrRegisteredStatusRequired = errors.New("registered status is required")
)

type ErrorMsg struct {
	Message string `json:"message"`
}
