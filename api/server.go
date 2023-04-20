package api

import (
	"github.com/aagolovanov/awesomeCodeService/domain"
	"github.com/aagolovanov/awesomeCodeService/util"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	router *http.ServeMux
	config util.Config
	domain *domain.Domain
	logg   *log.Logger
}

func NewServer(config util.Config, dom *domain.Domain, logger *log.Logger) *Server {
	router := http.NewServeMux()
	server := &Server{
		router: router,
		config: config,
		domain: dom,
		logg:   logger,
	}

	server.configureRoutes()

	return server
}

func (s *Server) Start() error {
	addr := ":" + strconv.Itoa(s.config.Port)

	return http.ListenAndServe(addr, s.router)
}
