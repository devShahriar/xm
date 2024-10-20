package http

import "github.com/devShahriar/xm/internal/common"

func (s *Server) RegisterRoutes() {
	s.Router.GET("/v1/companies/:id", s.Auth(Middleware(common.PermCommon, s.GetCompanyByID)))
	s.Router.PUT("/v1/companies/:id", s.Auth(Middleware(common.PermUpdateCompany, s.UpdateCompany)))
	s.Router.POST("/v1/companies", s.Auth(Middleware(common.PermCreateCompany, s.CreateCompany)))
	s.Router.DELETE("/v1/companies/:id", s.Auth(Middleware(common.PermDeleteCompany, s.DeleteCompany)))

	s.Router.POST("/v1/login", s.Login)
}
