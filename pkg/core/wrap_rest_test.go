package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsJSON(t *testing.T) {
	handler := func(r *RestRequest) interface{} {
		return map[string]string{
			"TestKey": "TestValue",
		}
	}
	rr := httptest.NewRecorder()
	httpHandler := WrapRest(handler)
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	httpHandler(rr, req)
	var respData map[string]string
	decodeErr := json.NewDecoder(rr.Body).Decode(&respData)
	if decodeErr != nil {
		t.Errorf("failed to decode response JSON: %v", decodeErr)
		return
	}
	if len(respData) != 1 || respData["TestKey"] != "TestValue" {
		t.Errorf("JSON return value didn't match returned value")
		return
	}
}

func TestDecodesJSONRequests(t *testing.T) {
	handler := func(r *RestRequest) interface{} {
		var reqData map[string]string
		decodeErr := r.DecodeJSON(&reqData)
		if decodeErr != nil {
			t.Errorf("failed to decode request JSON: %v", decodeErr)
			return nil
		}
		if len(reqData) != 1 || reqData["key"] != "value" {
			t.Errorf("JSON return value didn't match value from request")
			return nil
		}
		return nil
	}
	rr := httptest.NewRecorder()
	httpHandler := WrapRest(handler)
	req, err := http.NewRequest("POST", "", bytes.NewBuffer([]byte(`{"key": "value"}`)))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	httpHandler(rr, req)
}
