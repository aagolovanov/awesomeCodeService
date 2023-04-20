package api

import (
	"github.com/aagolovanov/awesomeCodeService/util"
	"net/http"
	"strconv"
)

type Server struct {
	router *http.ServeMux
	config util.Config
}

func NewServer(config util.Config) *Server {
	router := http.NewServeMux()
	configureRoutes(router)

	server := &Server{
		router: router,
		config: config,
	}

	return server
}

func (s *Server) Start() error {
	addr := ":" + strconv.Itoa(s.config.Port)

	return http.ListenAndServe(addr, s.router)
}
