package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/modules/data"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/modules/get"
	"github.com/gorilla/mux"
)

const (
	// default person id (returned on error)
	defaultPersonID = 0
)

// GetHandler is the HTTP handler for the "Get Person" endpoint
// In this simplified example we are assuming all possible errors are user errors and returning "bad request" HTTP 400
// or "not found" HTTP 404
// There are some programmer errors possible but hopefully these will be caught in testing.
type GetHandler struct {
}

// ServeHTTP implements http.Handler
func (h *GetHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// extract person id from request
	id, err := h.extractID(request)
	if err != nil {
		// output error
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// attempt get
	getter := get.Getter{}
	person, err := getter.Do(id)
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// happy path
	err = h.writeJSON(response, person)
	if err != nil {
		// this error should not happen but if it does there is nothing we can do to recover
		response.WriteHeader(http.StatusInternalServerError)
	}
}

// extract the person ID from the request
func (h *GetHandler) extractID(request *http.Request) (int, error) {
	// ID is part of the URL, so we extract it from there
	vars := mux.Vars(request)
	idAsString, exists := vars["id"]
	if !exists {
		// log and return error
		err := errors.New("[get] person id missing from request")
		logging.L.Warn(err.Error())
		return defaultPersonID, err
	}

	// convert ID to int
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		// log and return error
		err = fmt.Errorf("[get] failed to convert person id into a number. err: %s", err)
		logging.L.Error(err.Error())
		return defaultPersonID, err
	}

	return id, nil
}

// output the supplied person as JSON
func (h *GetHandler) writeJSON(writer io.Writer, person *data.Person) error {
	output := &getResponseFormat{
		ID:       person.ID,
		FullName: person.FullName,
		Phone:    person.Phone,
		Currency: person.Currency,
		Price:    person.Price,
	}

	// call to http.ResponseWriter.Write() will cause HTTP OK (200) to be output as well
	return json.NewEncoder(writer).Encode(output)
}

// the JSON response format
type getResponseFormat struct {
	ID       int     `json:"id"`
	FullName string  `json:"name"`
	Phone    string  `json:"phone"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
