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
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"

	// Adding the Postgres driver
	_ "github.com/lib/pq"
)

// openDbConnection is used to open a connection to a Postgres database
func openDbConnection() (*sql.DB, error) {
	// Validate the arguments provided
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName,
	)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database postgres://%s:%d/%s -> %w", dbHost, dbPort, dbName, err)
	}

	// Open up a connection to the database for compleate validation
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping database postgres://%s:%d/%s -> %w", dbHost, dbPort, dbName, err)
	}

	log.Printf("Successfully connected to database postgres://%s:%d/%s\n", dbHost, dbPort, dbName)
	return db, nil
}

// recordEventIntoDb inserts events into a database
func recordEventIntoDb(event Event) error {
	// Open a connection
	db, err := openDbConnection()
	if err != nil {
		return err
	}

	// Making sure that the connection is closed at the end
	defer db.Close()

	// Insert data
	query := fmt.Sprintf("INSERT INTO %s (ts, customer_id, system_name, billed_amount, msg) VALUES ($1, $2, $3, $4, $5)", tableName)
	_, err = db.Exec(query, event.Timestamp, event.CustomerId, event.SystemName, event.BilledAmount, event.Message)
	if err != nil {
		return fmt.Errorf("couldn't execute SQL statement <%s> -> %w", query, err)
	}

	return nil
}

// addSqlConditions will add a condition or a set of conditions based on the parameters that will be received
func addSqlConditions(values url.Values, query *string, paramName string) {
	paramValue, present := values[paramName]
	if present {
		action := "WHERE"
		if strings.Contains(*query, "WHERE") {
			action = "AND"
		}
		*query = *query + " " + action + " " + paramName + "='" + paramValue[0] + "'"
	}
}

// checkIfParamIsSupported will check if a given parameter is in the supported list
func checkIfParamIsSupported(givenParam string, supportedParams []string) bool {
	for _, supportedParam := range supportedParams {
		if givenParam == supportedParam {
			return true
		}
	}
	return false
}

// retrieveEventsFromDb gets the events from the database
func retrieveEventsFromDb(values url.Values) (Events, error) {
	// Open a connection
	db, err := openDbConnection()
	if err != nil {
		return Events{}, err
	}

	// Making sure that the connection is closed at the end
	defer db.Close()

	// Choose the SQL query based on the values recieved from the HTTP query
	query := fmt.Sprintf("SELECT ts, customer_id, system_name, billed_amount, msg FROM %s", tableName)
	if len(values) != 0 {
		for param := range values {
			if !checkIfParamIsSupported(param, supportedFieldValues) {
				return Events{}, fmt.Errorf("parameter %s is not supported", param)
			}
			addSqlConditions(values, &query, param)
		}
	}

	// Retrieve events
	rows, err := db.Query(query)
	if err != nil {
		return Events{}, err
	}
	defer rows.Close()
	var events Events
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Timestamp, &event.CustomerId, &event.SystemName, &event.BilledAmount, &event.Message)
		if err != nil {
			return Events{}, err
		}
		events = append(events, event)
	}

	return events, nil
}
