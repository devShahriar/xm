package http

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) GetCompanyByID(c echo.Context) error {
	id := c.Param("id")

	company, err := s.companyUC.GetCompanyByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Company not found")
	}

	return c.JSON(http.StatusOK, company)
}
