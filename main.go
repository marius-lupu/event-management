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
	"fmt"
	"log"
	"net/http"
	"os"
)

// getEnv is used to get the env variable or set a default one if it doesn't exists
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Record event
	http.HandleFunc("/event", recordEvent)

	// Retrieve events
	http.HandleFunc("/events", retrieveEvents)

	// Start listening
	log.Println("Event Management Service started ...")
	listenAddress := fmt.Sprintf("%s:%d", listenHost, listenPort)
	log.Println("Listening on ", listenAddress)
	http.ListenAndServe(listenAddress, nil)
}
