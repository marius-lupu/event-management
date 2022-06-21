SHELL=/bin/bash
GO111MODULE=on
BINARY_NAME=event-management
ADDRESS=http://localhost:3000
HEADERS=-H 'Content-Type: application/json' -H "Authorization: Basic ZXZlbnQ6cGFzc3dvcmQ="
CURL_COMMAND=curl $(HEADERS)
RECORD_COMMAND=$(CURL_COMMAND) -X POST $(ADDRESS)/event --data-binary

APP_USERNAME=event_user
DB_NAME=event_db

export GO111MODULE

.DEFAULT_GOAL := help

test:
	@clear
	export DB_HOST=localhost
	go test -v ./...
	echo
	go test -cover ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

tests: test

build:
	@mkdir -p bin
	go clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) .

run:
	@go mod tidy
	go run .

record:
	$(RECORD_COMMAND) '{"customerId": "abc-123456789", "systemName": "Earth", "message": "account was deleted"}'
	$(RECORD_COMMAND) '{"customerId": "def-922410780", "systemName": "Mars", "message": "resource modified by account"}'
	$(RECORD_COMMAND) '{"customerId": "abc-123456789", "systemName": "Mars", "billedAmount": 1000, "message": "new account created"}'
	$(RECORD_COMMAND) '{"customerId": "ghi-021429081", "systemName": "Phobos", "message": "account was deactivated"}'

# TODO: Currently, it doens't work with multiple field values
retrieve:
	$(CURL_COMMAND) $(ADDRESS)/events?customer_id=abc-123456789
	$(CURL_COMMAND) $(ADDRESS)/events?customer_id=ghi-021429081
	$(CURL_COMMAND) $(ADDRESS)/events?system_name=Phobos

start:
	docker-compose up --build -d

exec-db:
	docker-compose exec database psql -U $(APP_USERNAME) -d $(DB_NAME)

stop:
	docker-compose down -v

restart: stop start

status:
	docker-compose logs
	docker-compose ps

# WARNING: .ONESHELL not working for GNU Make =<3.81
.ONESHELL:
help:
	echo -e "
	================================================
	\t\tEvent Management
	================================================
	
	Commands available:
	\tmake test
	\tmake build
	\tmake run

	\tmake record
	\tmake retrieve
	\tmake exec-db

	\tmake status
	\tmake start
	\tmake restart
	\tmake stop
	"
