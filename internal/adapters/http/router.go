package http

func (s *Server) RegisterRoutes() {
	s.Router.GET("/v1/companies/:id", s.GetCompanyByID)
}
