package handlers

type Response struct {
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
}