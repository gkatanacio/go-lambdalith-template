package sample_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/go-lambdalith-template/internal/sample"
	mockssample "github.com/gkatanacio/go-lambdalith-template/mocks/sample"
)

func Test_Controller_Hello(t *testing.T) {
	// given
	ctx := context.Background()

	service := mockssample.NewMockService(t)
	service.EXPECT().Hello(ctx).Return("Hello there, John Doe!")

	controller := sample.NewController(service)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	// when
	controller.Hello(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"data\":\"Hello there, John Doe!\"}\n", rr.Body.String())
}

func Test_Controller_Echo(t *testing.T) {
	// given
	controller := sample.NewController(nil)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/echo", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"message":"Lorem ipsum dolor sit amet"}`))

	// when
	controller.Echo(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"data\":{\"message\":\"Lorem ipsum dolor sit amet\"}}\n", rr.Body.String())
}
