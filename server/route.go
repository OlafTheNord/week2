package server

import "net/http"

func SetupRoutes() *http.ServeMux {
	handler := NewHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees", handler.GetAll)
	mux.HandleFunc("GET /employee", handler.GetById)

	return mux
}
