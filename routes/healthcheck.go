package routes

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
	Status string `json:"status"`
}

// func NewHealthCheck(statusText string) *HealthCheck {
// 	return &HealthCheck{
// 		Status: statusText,
// 	}
// }

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	hc := HealthCheck{Status: "I am doing fine!"}
	jsonBytes, _ := json.Marshal(hc)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// func (h HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	jsonRepr, _ := json.Marshal(h)

// 	w.WriteHeader(200)
// 	w.Write(jsonRepr)
// }
