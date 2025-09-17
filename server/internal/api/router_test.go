package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRoute(t *testing.T) {
	router := NewRouter(nil)
	req := httptest.NewRequest("GET", "/api/health-check", nil)
	rw := httptest.NewRecorder()

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
}
