package api

import "net/http"

func configureRoutes(router *http.ServeMux) {
	router.HandleFunc("/api/v1/send", onlyPost(generateCode))
	router.HandleFunc("/api/v1/verify", onlyPost(verifyCode))
}
