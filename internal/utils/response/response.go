package response

import (
	"encoding/json"
	"net/http"
)
//Struct to show meaningful error
type Response struct{
	Status string
	Error string
}

const (
	StatusOK = "OK"
	StatusError = "Error" 
)

func WriteJson(w http.ResponseWriter, status int, data any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}