//go:build integration
// +build integration

package healthz

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const route = "/healthz"

// Integration test for Healthz method
func TestIntegrationHealthz(t *testing.T) {
	//Act
	response, _ := http.Get("http://localhost:8901/healthz")

	//Assert
	assert.Equal(t, "application/json; charset=UTF-8", response.Header.Get("Content-Type"))
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
