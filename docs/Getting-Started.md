# Instructions

## Prerequisites

- Docker and docker-compose to be installed in the system

## Config

The configuration will be in the `.env` file if it will be needed, otherwise the defaults will be used.

## Start

The following command will bring up all the microservices that will be needed

```bash
make start
```

To check the status, run the following:

```bash
make status
```

The last 3 lines of the ouput should look something like this:

```
NAME                        COMMAND                  SERVICE             STATUS              PORTS
event-management-event-1   "/event-management"     event               running             0.0.0.0:3000->8080/tcp
event-management-database-1      "docker-entrypoint.s…"   database                  running (healthy)   0.0.0.0:5432->5432/tcp
```

## Get the authorization string

This service is using the HTTP Authorization basic header.

```bash
APP_USERNAME=event
APP_PASSWORD=password
export APP_TOKEN=$(printf ${APP_USERNAME}:${APP_PASSWORD} | base64)
```

## Record events

```bash
curl \
    -H 'Content-Type: application/json' \
    -H "Authorization: Basic ${APP_TOKEN}" \
    -X POST http://localhost:3000/event \
    --data-binary \
    '{"customerId": "abc-123456789", "systemName": "Earth", "message": "account was deleted"}'
```

Other examples for a JSON payload:

```json
{"customerId": "def-922410780", "systemName": "Mars", "message": "resource modified by account"}
```

```json
{"customerId": "abc-123456789", "systemName": "Mars", "billedAmount": 1000, "message": "new account created"}
```

```json
{"customerId": "ghi-021429081", "systemName": "Phobos", "message": "account was deactivated"}
```

## Retrieve events

```bash
curl \
    -H 'Content-Type: application/json' \
    -H "Authorization: Basic ${APP_TOKEN}" \
    http://localhost:3000/events?customer_id=abc-123456789
```

Another endpoint that can be used as an example:

```
/events?system_name=Phobos
```

## Stop and Cleanup

To stop this service, run the following:

```bash
make stop
```
