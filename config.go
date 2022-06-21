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
	"time"
)

const (
	listenHost = "0.0.0.0"
	listenPort = 8080
	dbPort = 5432
)

var (
	appUsername          = getEnv("APP_USERNAME", "event")
	appPassword          = getEnv("APP_PASSWORD", "password")
	dbHost               = getEnv("DB_HOST", "postgres")
	dbUsername           = getEnv("DB_USERNAME", "event_user")
	dbPassword           = getEnv("DB_PASSWORD", "password")
	dbName               = getEnv("DB_NAME", "event_db")
	tableName            = getEnv("DB_TABLE_NAME", "events")
	supportedFieldValues = []string{"customer_id", "system_name"}
)

type Event struct {
	Id           int64     `json:"id,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	SystemName   string    `json:"systemName"`
	CustomerId   string    `json:"customerId"`
	BilledAmount *int64    `json:"billedAmount,omitempty"`
	Message      string    `json:"message"`
}

type Events []Event
