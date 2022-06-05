package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/service"
	"github.com/sirupsen/logrus"
)

type PoliclinicHandler struct {
	service service.Patient
}

func NewPoliclinicHandler(service service.Patient) *PoliclinicHandler {
	return &PoliclinicHandler{
		service: service,
	}
}

func (h *PoliclinicHandler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("GET params were:", r.URL.Query())
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		e, err := h.service.GetDoctors()
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendDoctorResponse(w, http.StatusOK, e)
	}
}

func (h *PoliclinicHandler) GetRecords(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("GET params were:", r.URL.Query())
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		e, err := h.service.GetRecords()
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendRecordsResponse(w, http.StatusOK, e)
	}
}

func (h *PoliclinicHandler) GetRecordOfPatient(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		patientId, err := strconv.Atoi(id)
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          fmt.Errorf("id - число").Error(),
				HTTPStatusCode: http.StatusMethodNotAllowed,
			})
		}
		e, err := h.service.GetRecordOfPatient(patientId)
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendRecordsResponse(w, http.StatusOK, e)
	}
}

func (h *PoliclinicHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient labsMed.Patient
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusBadRequest,
			})
			logrus.Println("Error decoding", err.Error(), patient)
			return
		}
		e, err := h.service.CreatePatient(patient)
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendResponse(w, http.StatusOK, e)
	}
}

func (h *PoliclinicHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor labsMed.Doctor
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusBadRequest,
			})
			logrus.Println("Error decoding", err.Error(), doctor)
			return
		}
		e, err := h.service.CreateDoctor(doctor)
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendResponse(w, http.StatusOK, e)
	}
}

func (h *PoliclinicHandler) CreatePatientRecord(w http.ResponseWriter, r *http.Request) {
	var record labsMed.Record
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusBadRequest,
			})
			logrus.Println("Error decoding", err.Error(), record)
			return
		}
		e, err := h.service.CreatePatientRecord(record)
		if err != nil {
			sendErrorResponse(w, r, &ErrorModel{
				Error:          err.Error(),
				HTTPStatusCode: http.StatusServiceUnavailable,
			})
			logrus.Println(err)
			return
		}
		sendResponse(w, http.StatusOK, e)
	}
}

func sendResponse(w http.ResponseWriter, s int, event int) {
	w.WriteHeader(s)
	data, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		logrus.Fatal(err)
	}
	w.Write(data)
}

func sendErrorResponse(w http.ResponseWriter, r *http.Request, e *ErrorModel) {
	w.WriteHeader(e.HTTPStatusCode)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(e)
	w.Write(reqBodyBytes.Bytes())

}

type ErrorModel struct {
	Error          string
	HTTPStatusCode int
}

func sendDoctorResponse(w http.ResponseWriter, s int, event []labsMed.Doctor) {
	w.WriteHeader(s)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(event)
	data, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		logrus.Fatal(err)
	}
	w.Write(data)
}

func sendRecordsResponse(w http.ResponseWriter, s int, event []labsMed.Record) {
	w.WriteHeader(s)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(event)
	data, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		logrus.Fatal(err)
	}
	w.Write(data)
}
