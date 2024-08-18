package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse represents a JSON response.
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SendJSONResponse sends a JSON response.
func SendJSONResponse(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JSONResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// SendErrorResponse sends an error response.
func SendErrorResponse(w http.ResponseWriter, code int, msg string) {
	SendJSONResponse(w, code, msg, nil)
}

// SendSuccessResponse sends a success response.
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	SendJSONResponse(w, http.StatusOK, "success", data)
}
