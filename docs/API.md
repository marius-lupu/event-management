# API

## HEADERS

The headers that will be required will be:

- `Content-Type: application/json`
- `Authorization: Basic <TOKEN>`, where `<TOKEN>` is the encoded username:password

## GET /events

This will get the events from the system.
Querying the recorded event data is accepted also by using some field values.
There are 2 query parameter that are currently accepted:

- customer_id
- system_name

Examples:

- `/events`
- `/events?customer_id=abc-123456789`
- `/events?system_name=Phobos`

## POST /event

This endpoint is used to add a new event.

Payload info:

| Parameter Name | Description | Accepted Type | Mandatory | Example |
|---|---|---|---|---|
| `customerId` | This is the customer unique ID | String | Yes | `abc-123456789` |
| `systemName` | This will be the name of the system where the request originates | String | Yes | `Pluto` |
| `message` | This is the text message | String | Yes | `resource modified by account` |
| `billedAmount` | This is the amount that the customer will be billed | Integer | No | `1000` |

Examples:

```json
{"customerId": "abc-123456789", "systemName": "Earth", "message": "account was deleted"}
{"customerId": "def-922410780", "systemName": "Mars", "message": "resource modified by account"}
{"customerId": "abc-123456789", "systemName": "Mars", "billedAmount": 1000, "message": "new account created"}
{"customerId": "ghi-021429081", "systemName": "Phobos", "message": "account was deactivated"}
```