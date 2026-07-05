package httputil

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Problem struct {
	Status  int    `json:"status"`
	Title   string `json:"title"`
	Detail  string `json:"detail,omitempty"`
	Code    string `json:"code,omitempty"`
}

type AppError struct {
	Status int
	Code   string
	Title  string
	Err    error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Title
}

func NewAppError(status int, code, title string, err error) *AppError {
	return &AppError{Status: status, Code: code, Title: title, Err: err}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("write json response", "error", err)
	}
}

func WriteError(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		WriteJSON(w, appErr.Status, Problem{
			Status: appErr.Status,
			Title:  appErr.Title,
			Detail: appErr.Err.Error(),
			Code:   appErr.Code,
		})
		return
	}

	WriteJSON(w, http.StatusInternalServerError, Problem{
		Status: http.StatusInternalServerError,
		Title:  "internal server error",
		Detail: err.Error(),
	})
}

func HealthHandler(checks map[string]func() error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		result := make(map[string]string, len(checks))

		for name, check := range checks {
			if err := check(); err != nil {
				result[name] = "down"
				status = http.StatusServiceUnavailable
			} else {
				result[name] = "up"
			}
		}

		WriteJSON(w, status, map[string]any{
			"status": map[bool]string{true: "ok", false: "degraded"}[status == http.StatusOK],
			"checks": result,
		})
	}
}
