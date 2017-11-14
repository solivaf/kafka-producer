package cmd

import (
	"github.com/solivaf/kafka-producer/cmd/handlers"
	"net/http"
	"github.com/gorilla/mux"
)

func NewServer() {

	h := &handlers.HealthCheckHandler{}
	mp := &handlers.ProducerHandler{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{topic}", mp.MessageProducer).Methods("POST")
	router.HandleFunc("/health", h.HealthCheck).Methods("GET")

	http.ListenAndServe(":8080", router)
}
