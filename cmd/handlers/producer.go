package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
)

type ProducerHandler struct {}

func (p *ProducerHandler) MessageProducer(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(Response{Message: "Error while reading body"})
		return
	}
	vars := mux.Vars(request)
	topic := vars["topic"]



	writer.WriteHeader(201)
}