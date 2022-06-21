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
	"net/http"
	"net/http/httptest"
	"testing"
)

// sendRetrieveRequest sends a request to record an event
func sendRetrieveRequest(t *testing.T, method string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, "/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth(appUsername, appPassword)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(retrieveEvents)
	handler.ServeHTTP(rr, req)

	return rr
}

// TestRetrieveEvents will test the retrieve events endpoint
func TestRetrieveEvents(t *testing.T) {
	// Positive Case
	rr := sendRetrieveRequest(t, http.MethodGet)
	if rr.Code != http.StatusOK {
		t.Fatal("GET request didn't return 200")
	}

	// Negative Case
	rr = sendRetrieveRequest(t, http.MethodDelete)
	checkResponse(t, rr, http.StatusMethodNotAllowed, `DELETE method not accepted for retrieving events, please use GET`)
}
