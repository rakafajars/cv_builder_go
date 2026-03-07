package pkg

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

type Response struct {
	Status       string `json:"status"`
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message,omitempty"`
	Data         any    `json:"data,omitempty"`
	Error        string `json:"error,omitempty"`
}

type ResponsePaginated struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Error        string         `json:"error,omitempty"`
	Meta         PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page      int    `json:"page" example:"1"`
	Limit     int    `json:"limit" example:"10"`
	Total     int    `json:"total" example:"100"`
	TotalPage int    `json:"total_pages" example:"10"`
	Filter    string `json:"filter" example:"nama=triady"`
	Sort      string `json:"sort" example:"-id"`
}

func Success(w http.ResponseWriter, message string, data any) error {
	return writeJSON(w, http.StatusOK, Response{
		Status:       "Success",
		ResponseCode: http.StatusOK,
		Message:      message,
		Data:         data,
	})
}

func SuccessPagination(w http.ResponseWriter, message string, data any, meta PaginationMeta) error {
	return writeJSON(w, http.StatusOK, ResponsePaginated{
		Status:       "Success",
		ResponseCode: http.StatusOK,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Created(w http.ResponseWriter, message string, data any) error {
	return writeJSON(w, http.StatusCreated, Response{
		Status:       "Created",
		ResponseCode: http.StatusCreated,
		Message:      message,
		Data:         data,
	})
}

func BadRequest(w http.ResponseWriter, message string, err string) error {
	return writeJSON(w, http.StatusBadRequest, Response{
		Status:       "Error Bad Request",
		ResponseCode: http.StatusBadRequest,
		Message:      message,
		Error:        err,
	})
}

func NotFound(w http.ResponseWriter, message string, err string) error {
	return writeJSON(w, http.StatusNotFound, Response{
		Status:       "Error Not Found",
		ResponseCode: http.StatusNotFound,
		Message:      message,
		Error:        err,
	})
}

func NotFoundPagination(w http.ResponseWriter, message string, data any, meta PaginationMeta) error {
	return writeJSON(w, http.StatusNotFound, ResponsePaginated{
		Status:       "Not Found",
		ResponseCode: http.StatusNotFound,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Unauthorized(w http.ResponseWriter, message string, err string) error {
	return writeJSON(w, http.StatusUnauthorized, Response{
		Status:       "Error Unauthorized",
		ResponseCode: http.StatusUnauthorized,
		Message:      message,
		Error:        err,
	})
}

func InternalServerError(w http.ResponseWriter, message string, err string) error {
	return writeJSON(w, http.StatusInternalServerError, Response{
		Status:       "Internal Server Error",
		ResponseCode: http.StatusInternalServerError,
		Message:      message,
		Error:        err,
	})
}
