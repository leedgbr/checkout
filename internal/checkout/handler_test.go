package checkout

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	type test struct {
		name             string
		request          string
		stub             service
		expectedStatus   int
		expectedResponse string
	}

	tests := []test{
		{"success", `{"sku": ["120P90"]}`, &stubService{}, http.StatusOK, `{"total":1234}`},
		{"business error", `{"sku": ["123456"]}`, &stubAlwaysErrorService{}, http.StatusUnprocessableEntity, `{"message":"business error"}`},
		{"bad request", `rubbish`, &stubService{}, http.StatusBadRequest, `{"message":"error parsing request"}`},
	}

	for _, tc := range tests {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(newHandler(tc.stub))
		handler.ServeHTTP(rr, newRequest(t, tc.request))
		assert.Equal(t, tc.expectedStatus, rr.Code, "unexpected http status")
		assert.JSONEq(t, tc.expectedResponse, rr.Body.String(), "unexpected response")
	}
}

func TestHandlerErrorEncodingResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	encode(rr, math.Inf(1))
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "system error\n", rr.Body.String())
}

func newRequest(t *testing.T, payload string) *http.Request {
	req, err := http.NewRequest("GET", "/", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

type stubService struct {
}

func (s stubService) checkout(sku []string) (*result, error) {
	return &result{TotalInCents: 1234}, nil
}

type stubAlwaysErrorService struct {
}

func (s stubAlwaysErrorService) checkout(sku []string) (*result, error) {
	return nil, errors.New("business error")
}
