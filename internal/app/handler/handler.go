package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mux-api-boilerplate/internal/app/domain"
	"mux-api-boilerplate/internal/app/models"
	"net/http"
)

// Handler is the main adapter to the API. It translates the upcoming information from the external port
// to the application language, starting the request flow.
type Handler struct {
	service domain.Service
}

// NewHandler returns an instance of our handler
func NewHandler(service domain.Service) Handler {
	return Handler{
		service: service,
	}
}

// Find returns a single object matching an ID
func (h Handler) Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingRequiredID)
		return
	}
	resp, err := h.service.GetById(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

// Get returns all objects
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

// Post creates an object
func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	var obj models.SampleModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidRequestPayload)
		return
	}
	defer r.Body.Close()

	resp, err := h.service.Create(obj)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

// Put updates an entire object
func (h Handler) Put(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingRequiredID)
		return
	}

	var obj models.SampleModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidRequestPayload)
		return
	}
	defer r.Body.Close()

	err := h.service.Update(id, obj)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// Patch updates an attribute from an object
func (h Handler) Patch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingRequiredID)
		return
	}

	attribute, ok := params["attribute"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingAttributeName)
		return
	}

	value, ok := params["value"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingAttributeValue)
		return
	}

	err := h.service.UpdateAttribute(id, attribute, value)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// Delete removes the object
func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, ErrorMissingRequiredID)
		return
	}
	err := h.service.Remove(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
