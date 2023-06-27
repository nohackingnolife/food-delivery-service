package response

import "net/http"

type APIError struct {
	Message string `json:"message"`
}

func NewAPIError(message string) *APIError {
	return &APIError{
		Message: message,
	}
}

func SendNotFoundError(w http.ResponseWriter, err error) {
	sendJSON(w, http.StatusNotFound, Basic{
		Message: err.Error(),
	})
}

func SendMethodNotAllowedError(w http.ResponseWriter, err error) {
	sendJSON(w, http.StatusMethodNotAllowed, Basic{
		Message: err.Error(),
	})
}

func SendBadRequestError(w http.ResponseWriter, err error) {
	sendJSON(w, http.StatusBadRequest, Basic{
		Message: err.Error(),
	})
}

func SendInternalServerError(w http.ResponseWriter, err error) {
	sendJSON(w, http.StatusInternalServerError, Basic{
		Message: err.Error(),
	})
}
