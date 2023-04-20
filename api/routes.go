package api

func (s *Server) configureRoutes() {
	s.router.HandleFunc("/api/v1/send", onlyPost(s.generateCode))
	s.router.HandleFunc("/api/v1/verify", onlyPost(s.verifyCode))
}
