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
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// recordEvent is a handler that will be used to record events
func recordEvent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receiving a hit in the %s endpoint\n", r.URL.Path)

	// Accept only the POST method
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("ERROR: Method %s used instead of POST\n", r.Method)
		fmt.Fprintf(w, "%s method not accepted for recording an event, please use POST", r.Method)
		return
	}

	// Verify credentials
	err := verifyCredentials(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("ERROR:", err)
		fmt.Fprintf(w, "%s", err)
		return
	}

	// Get the event data
	var event Event
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Coudln't get the binary data from the body ->", err)
		fmt.Fprintf(w, "Couldn't process the request payload")
		return
	}
	err = json.Unmarshal(bb, &event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Coudln't unmarshal data into the Event struct ->", err)
		fmt.Fprintf(w, "Payload request not supported")
		return
	}

	// Add timestamp if empty
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Record event data
	err = recordEventIntoDb(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR: Coudln't record event into database ->", err)
		fmt.Fprintf(w, "Unexpected error occured while trying to record the event into the database")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Event recorded successfully")
}
