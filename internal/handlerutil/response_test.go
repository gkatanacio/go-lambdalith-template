package handlerutil_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gkatanacio/go-lambdalith-template/internal/handlerutil"
	mockshttp "github.com/gkatanacio/go-lambdalith-template/mocks/http"
)

func Test_JsonResponse(t *testing.T) {
	testCases := map[string]struct {
		data     any
		wantBody string
	}{
		"map object": {
			data: map[string]string{
				"firstName": "John",
				"lastName":  "Doe",
			},
			wantBody: `{"firstName":"John","lastName":"Doe"}`,
		},
		"array": {
			data:     []bool{true, false, true},
			wantBody: "[true,false,true]",
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			testStatusCode := http.StatusOK

			writer := mockshttp.NewMockResponseWriter(t)
			writer.EXPECT().Header().Return(make(http.Header))
			writer.EXPECT().WriteHeader(testStatusCode)
			writer.EXPECT().Write([]byte(tc.wantBody+"\n")).Return(0, nil)

			handlerutil.JsonResponse(writer, testStatusCode, tc.data)
		})
	}
}

func Test_ErrorResponse(t *testing.T) {
	testCases := map[string]struct {
		err        error
		wantStatus int
		wantBody   string
	}{
		"generic error": {
			err:        errors.New("something went wrong"),
			wantBody:   `{"error":"something went wrong"}`,
			wantStatus: http.StatusInternalServerError,
		},
		"bad request": {
			err:        handlerutil.BadRequest("invalid input"),
			wantBody:   `{"error":"invalid input"}`,
			wantStatus: http.StatusBadRequest,
		},
		"not found": {
			err:        handlerutil.NotFound("not found"),
			wantBody:   `{"error":"not found"}`,
			wantStatus: http.StatusNotFound,
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			writer := mockshttp.NewMockResponseWriter(t)
			writer.EXPECT().Header().Return(make(http.Header))
			writer.EXPECT().WriteHeader(tc.wantStatus)
			writer.EXPECT().Write([]byte(tc.wantBody+"\n")).Return(0, nil)

			handlerutil.ErrorResponse(writer, tc.err)
		})
	}
}
