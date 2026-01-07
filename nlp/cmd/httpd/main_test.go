package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Basic approach to testing the server
func Test_healthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	healthHandler(w, r)
	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
