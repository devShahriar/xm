package http

import (
	"github.com/devShahriar/xm/internal/usecase"
	"github.com/labstack/echo"
)

type Server struct {
	Router    *echo.Echo
	companyUC *usecase.CompanyUsecase
}

func NewServer(companyUC *usecase.CompanyUsecase) *Server {
	echoServer := echo.New()
	return &Server{
		Router:    echoServer,
		companyUC: companyUC,
	}
}

func (s *Server) Start() {
	s.Router.Logger.Fatal(s.Router.Start(":8090"))
}
