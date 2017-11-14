package handlers

import (
	"net/http"
)

type HealthCheckHandler struct {}

func (h *HealthCheckHandler) HealthCheck(writer http.ResponseWriter, request *http.Request) {
}