package handler

import (
	"net/http"

	"github.com/kirktriplefive/labsMed/pkg/service"
)

type HandlerInterfacePatient interface {
	CreatePatient(w http.ResponseWriter, r *http.Request)
	CreatePatientRecord(w http.ResponseWriter, r *http.Request)
	CreateDoctor(w http.ResponseWriter, r *http.Request)
	GetDoctors(w http.ResponseWriter, r *http.Request)
	GetRecords(w http.ResponseWriter, r *http.Request)
	GetRecordOfPatient(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	HandlerInterfacePatient
}

func NewHandlerPatient(service *service.ServicePatient) *Handler {
	return &Handler{&PoliclinicHandler{*service}}
}
