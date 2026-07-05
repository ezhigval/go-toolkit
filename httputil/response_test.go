package httputil

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteError_AppError(t *testing.T) {
	rec := httptest.NewRecorder()
	err := NewAppError(404, "NOT_FOUND", "link not found", errors.New("db miss"))

	WriteError(rec, err)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var body Problem
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "NOT_FOUND", body.Code)
	assert.Equal(t, "db miss", body.Detail)
}

func TestWriteError_GenericDoesNotLeak(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteError(rec, errors.New("secret connection string"))

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var body Problem
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, "internal server error", body.Title)
	assert.NotContains(t, body.Detail, "secret")
}
