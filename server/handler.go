package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"week2/service"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err, employees := service.EmployeeGetAll("src/employees.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		log.Fatal(err)
	}
}
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
	}
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, emp := service.EmployeeGetById("src/employees.json", ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if emp.Name == "" {
		http.Error(w, "employee not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(emp)
	if err != nil {
		log.Fatal(err)
	}

}
