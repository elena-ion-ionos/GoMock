package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_PostNews(t *testing.T) {
	test_cases := []struct {
		name               string
		expectedStatusCode int
	}{
		{name: "not implemented", expectedStatusCode: http.StatusNotImplemented},
	}
	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			//here is the logic for the subtest
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/", nil)
			
			// Get the handler function and call it
			handler := PostNewsHandler()
			handler(recorder, request)
			
			// Assert the expected status code
			if recorder.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, recorder.Code)
			}
		})
	}
}
