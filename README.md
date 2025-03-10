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
#### ShortenedAPI
|Path|Method|Description|
|----|------|-----------|
|`/rest/api`|POST|Create ShortenedAPI from JSON|
|`/rest/api/:id`|DELETE|Delete ShortenedAPI by ID|
|`/api/:id`|any|Send request to the ShortenedAPI with the ID|

#### OutgoingRequestConfig
|Path|Method|Description|
|----|------|-----------|
|`/rest/configs`|POST|Create OutgoingRequestConfig from JSON|
|`/rest/configs/:id`|GET|Get OutgoingRequestConfig by ID|
|`/rest/configs/?apiID=API_ID`|GET|Get OutgoingRequestConfig by ShortenedAPI ID|
|`/rest/configs`|PUT|Update OutgoingRequestConfig from JSON|
|`/rest/configs/:id`|DELETE|Delete OutgoingRequestConfig by ID|

#### ShorteningRule
|Path|Method|Description|
|----|------|-----------|
|`/rest/rules`|POST|Create ShorteningRule from JSON|
|`/rest/rules/:id`|GET|Get ShorteningRule by ID|
|`/rest/rules/?apiID=API_ID`|GET|Get all ShorteningRules by ShortenedAPI ID|
|`/rest/rules`|PUT|Update ShorteningRule from JSON|
|`/rest/rules/:id`|DELETE|Delete ShorteningRule by ID|

#### OutgoingRequestHeader
|Path|Method|Description|
|----|------|-----------|
|`/rest/headers`|POST|Create OutgoingRequestHeader from JSON|
|`/rest/headers/:id`|GET|Get OutgoingRequestHeader by ID|
|`/rest/headers/?configID=CONFIG_ID`|GET|Get all OutgoingRequestHeaders by OutgoingRequestConfig ID|
|`/rest/headers`|PUT|Update OutgoingRequestHeader from JSON|
|`/rest/headers/:id`|DELETE|Delete OutgoingRequestHeader by ID|

#### OutgoingRequestParam
|Path|Method|Description|
|----|------|-----------|
|`/rest/params`|POST|Create OutgoingRequestParam from JSON|
|`/rest/params/:id`|GET|Get OutgoingRequestParam by ID|
|`/rest/params/?configID=CONFIG_ID`|GET|Get all OutgoingRequestParams by OutgoingRequestConfig ID|
|`/rest/params`|PUT|Update OutgoingRequestParam from JSON|
|`/rest/params/:id`|DELETE|Delete OutgoingRequestParam by ID|
