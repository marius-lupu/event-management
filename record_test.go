/*
	Copyright (c) 2022 Marius Lupu

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// sendRecordRequest sends a request to record an event
func sendRecordRequest(t *testing.T, method string, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, "/event", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth(appUsername, appPassword)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(recordEvent)
	handler.ServeHTTP(rr, req)

	return rr
}

// TestRecordEvents will test the record event endpoint
func TestRecordEvents(t *testing.T) {
	// Positive Case
	event := Event{Timestamp: time.Now(), CustomerId: "abc-123456789", Message: "Test"}
	eventJson, err := json.Marshal(event)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	rr := sendRecordRequest(t, http.MethodPost, string(eventJson))
	checkResponse(t, rr, http.StatusOK, `Event recorded successfully`)

	// Negative Cases
	rr = sendRecordRequest(t, http.MethodGet, "")
	checkResponse(t, rr, http.StatusMethodNotAllowed, `GET method not accepted for recording an event, please use POST`)

	rr = sendRecordRequest(t, http.MethodPost, "Blabla")
	checkResponse(t, rr, http.StatusBadRequest, `Payload request not supported`)

	rr = sendRecordRequest(t, http.MethodPost, `{"timestamp": "2022-06-07T11:23:23.741657+03:00", "customerId": true, "message": "Test"}`)
	checkResponse(t, rr, http.StatusBadRequest, `Payload request not supported`)
}
