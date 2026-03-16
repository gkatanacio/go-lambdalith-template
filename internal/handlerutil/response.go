package handlerutil

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(data)
}

func DataResponse(w http.ResponseWriter, statusCode int, data any) {
	JsonResponse(w, statusCode, map[string]any{"data": data})
}

type errorResponseBody struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, err error) {
	var status int
	var httpErr HttpError
	switch {
	case errors.As(err, &httpErr):
		status = httpErr.StatusCode()
	default:
		status = http.StatusInternalServerError
	}

	body := &errorResponseBody{
		Error: err.Error(),
	}

	JsonResponse(w, status, body)
}
