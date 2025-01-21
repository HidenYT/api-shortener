# API Shortener

A proxy-microservice for shortening API responses using JSONPath language.

Can be helpful for systems with a limited storage size or low computational capabilities

## Main features:
- CRUD operations for API to be shortened
- Querying target APIs and selection of values from them using JSONPath language
- Token-based authentication
- Limited number of requests to a single ShortenedAPI that can be processed at 
the same time (in order to avoid loops when microservice recursively sends 
requests to itself)

## Starting the microservice
Duplicate the .env.example file, rename the copy to `.env` and fill in the values:

| ENV key | Description |
| --------|-------------|
|`API_KEY`|Strong secret key to authenticate in the microservice|
|`LOOP_LIMITER_MAX_REQUESTS`|The maximum number of requests to one ShortenedAPI that can be processed at the same time. If the number is `1` only one request to the API can be processed at the same time|
|`DB_USER`|Database user name|
|`DB_PASSWORD`|Database user password|
|`DB_NAME`|Database name|
|`OUTGOING_REQUEST_CLIENT_TIMEOUT`|Timeout for the requests to the APIs|
|`OUTGOING_REQUEST_CLIENT_RETRIES_COUNT`|The number of retries for the requests to the APIs|

The following command will automatically migrate your database and start the microservice:
```
go run api-shortener/main
```

## API 
### Authentication
For every request to the API of the microservice the client should pass a query
parameter `token` with the token from `API_KEY` in the `.env` file:
```
curl http://localhost:8080/api/63?token=MY_TOKEN_FROM_API_KEY
```
### Description
|Path|Method|Description|
|----|------|-----------|
|`/rest/:id`|GET|Get ShortenedAPI by ID|
|`/rest`|POST|Create ShortenedAPI from JSON|
|`/rest`|PUT|Update ShortenedAPI from JSON|
|`/rest/:id`|DELETE|Delete ShortenedAPI by ID|
|`/api/:id`|any|Send request to the ShortenedAPI with the ID|

### Examples
#### Create ShortenedAPI 
<details>
    <summary>Request</summary>

```
curl --location 'http://localhost:8080/rest?token=MY_TOKEN_FROM_API_KEY' \
--request POST \
--header 'Content-Type: application/json' \
--data '{
    "config": {
        "url": "https://timeapi.io/api/time/current/zone",
        "method": "GET",
        "params": [
            {
                "name": "timeZone",
                "value": "Europe/Amsterdam"
            }
        ],
        "headers": [
            {
                "name": "Accept",
                "value": "application/json"
            }
        ]
    },
    "rules": [
        {
            "field_name": "datetime",
            "field_value_query": "$.dateTime"
        }
    ]
}'
```
</details>

<details>
    <summary>Response body</summary>
    
```
{
    "id": 1,
    "config": {
        "id": 2,
        "url": "https://timeapi.io/api/time/current/zone",
        "method": "GET",
        "headers": [
            {
                "id": 3,
                "name": "Accept",
                "value": "application/json"
            }
        ],
        "params": [
            {
                "id": 4,
                "name": "timeZone",
                "value": "Europe/Amsterdam"
            }
        ],
        "body": ""
    },
    "rules": [
        {
            "id": 5,
            "field_name": "datetime",
            "field_value_query": "$.dateTime"
        }
    ]
}
```
</details>

#### Get ShortenedAPI
<details>
    <summary>Request</summary>

```
curl --location 'http://localhost:8080/rest/66?token=MY_TOKEN_FROM_API_KEY' --request GET
```
</details>
<details>
    <summary>Response body</summary>
    
```
{
    "id": 66,
    "config": {
        "id": 2,
        "url": "https://timeapi.io/api/time/current/zone",
        "method": "GET",
        "headers": [],
        "params": [
            {
                "id": 4,
                "name": "timeZone",
                "value": "Europe/Amsterdam"
            }
        ],
        "body": ""
    },
    "rules": [
        {
            "id": 5,
            "field_name": "datetime",
            "field_value_query": "$.dateTime"
        }
    ]
}
```
</details>

#### Update ShortenedAPI 
<details>
    <summary>Request</summary>

```
curl --location --request PUT 'http://localhost:8080/rest?token=MY_TOKEN_FROM_API_KEY' \
--request PUT \
--header 'Content-Type: application/json' \
--data '{
    "id": 69,
    "config": {
        "id": 72,
        "url": "https://timeapi.io/api/time/current/zone",
        "method": "GET",
        "headers": [],
        "params": [],
        "body": ""
    },
    "rules": [
        {
            "id": 66,
            "field_name": "datetime",
            "field_value_query": "$.dateTime"
        }
    ]
}'
```
</details>

<details>
    <summary>Response body</summary>

```
{
    "id": 69,
    "config": {
        "id": 72,
        "url": "https://timeapi.io/api/time/current/zone",
        "method": "GET",
        "headers": [],
        "params": [],
        "body": ""
    },
    "rules": [
        {
            "id": 66,
            "field_name": "datetime",
            "field_value_query": "$.dateTime"
        }
    ]
}
```
</details>

#### Delete ShortenedAPI
<details>
    <summary>Request</summary>

```
curl --location 'http://localhost:8080/rest/66?token=MY_TOKEN_FROM_API_KEY' --request DELETE
```
</details>
<details>
    <summary>Response body</summary>
    
```
{}
```
</details>