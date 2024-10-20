package http

import (
	"net/http"

	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func (s *Server) GetCompanyByID(c echo.Context) error {
	id := c.Param("id")

	company, err := s.companyUC.GetCompanyByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Company not found")
	}

	return c.JSON(http.StatusOK, company)
}

func (s *Server) CreateCompany(c echo.Context) error {
	// Bind the incoming JSON request to the Company entity
	company := new(entity.Company)
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	// Generate a new UUID for the company
	company.ID = uuid.New()

	// Call the use case to handle the business logic
	err := s.companyUC.CreateCompany(c.Request().Context(), company)
	if err != nil {
		log.Errorf("error while creating company %v", err)
		return c.JSON(http.StatusInternalServerError, common.ErrorMsg{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, company)
}

func (s *Server) UpdateCompany(c echo.Context) error {
	id := c.Param("id")
	companyID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid company ID")
	}

	// Bind the request body to the Company entity
	company := new(entity.Company)
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input data")
	}

	company.ID = companyID

	// Call the use case to handle the update logic
	updatedCompany, err := s.companyUC.UpdateCompany(c.Request().Context(), company)
	if err != nil {
		if err == common.ErrCompanyNotFound {
			return c.JSON(http.StatusNotFound, "Company not found")
		}
		return c.JSON(http.StatusInternalServerError, common.ErrorMsg{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, updatedCompany)
}

func (s *Server) DeleteCompany(c echo.Context) error {
	id := c.Param("id")
	companyID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid company ID")
	}

	// Call the use case to handle the deletion logic
	err = s.companyUC.DeleteCompany(c.Request().Context(), companyID)
	if err != nil {
		if err == common.ErrCompanyNotFound {
			return c.JSON(http.StatusNotFound, "Company not found")
		}
		return c.JSON(http.StatusInternalServerError, "Failed to delete company")
	}

	return c.NoContent(http.StatusNoContent) // Return 204 No Content if successful
}
