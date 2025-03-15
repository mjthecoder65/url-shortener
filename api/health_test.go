package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	require.NoError(t, err)

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusOK, res.Code)
}
