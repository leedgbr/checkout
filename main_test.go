package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"company.com/checkout/internal/checkout"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	type test struct {
		name             string
		request          string
		expectedStatus   int
		expectedResponse string
	}

	tests := []test{
		{"success", `{"sku": ["120P90"]}`, http.StatusOK, `{"total":4999}`},
		{"error", `{"sku": ["123456"]}`, http.StatusUnprocessableEntity, `{"message": "item not found"}`},
	}

	app := &checkout.App{}
	app.Init()

	for _, tc := range tests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/checkout", strings.NewReader(tc.request))
		req.Header.Set("Content-Type", "application/json")
		app.Router.ServeHTTP(rr, req)
		assert.Equalf(t, rr.Code, tc.expectedStatus, "unexpected http status")
		assert.JSONEq(t, tc.expectedResponse, rr.Body.String(), "unexpected response")
	}
}
