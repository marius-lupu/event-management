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
	"fmt"
	"log"
	"net/http"
)

// retrieveEvents is a handler that will be used to retrieve events
func retrieveEvents(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receiving a hit in the %s endpoint\n", r.URL.Path)

	// Accept only the GET method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("ERROR: Method %s used instead of GET\n", r.Method)
		fmt.Fprintf(w, "%s method not accepted for retrieving events, please use GET", r.Method)
		return
	}

	// Verify credentials
	err := verifyCredentials(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("ERROR: %s\n", err)
		fmt.Fprintf(w, "%s", err)
		return
	}

	// Retrieve events
	events, err := retrieveEventsFromDb(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERROR: %s\n", err)
		fmt.Fprintf(w, "Internal error occured, couldn't read events from the database")
		return
	}

	// No events
	if len(events) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "No events recorded")
		return
	}

	jsonResponse, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR: Couldn't marshal the events object ->", err)
		fmt.Fprintf(w, "Internal error occured, couldn't marshal the events object")
		return
	}

	// Return
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
