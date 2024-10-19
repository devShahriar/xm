package http

func (s *Server) RegisterRoutes() {
	s.Router.GET("/v1/companies/:id", s.GetCompanyByID)
	s.Router.PUT("/v1/companies/:id", s.UpdateCompany)
	s.Router.POST("/v1/companies", s.CreateCompany)
	s.Router.DELETE("/v1/companies/:id", s.DeleteCompany)
}
