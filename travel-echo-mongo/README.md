# Travel API using Go + echo + mgo

### Frameworks | Tools
1. [echo](https://echo.labstack.com/) - Go web framework | High performance, extensible, minimalist
2. [mgo](https://gopkg.in/mgo.v2) - Go Mongo DB Driver
3. [jwt-go](https://github.com/dgrijalva/jwt-go) - Go implementation of JSON Web Tokens [JWT](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html)

## Running (DOCKER)
Run `docker-compose up -d`

## Running (OFFLINE)
Make sure there is local mongoDB running in machine.

Run in command line go build server.go

## Signup
User signup

Retrieves credentials from the body and validate against database.
For invalid email or password, `send 400 - Bad Request` response.
For valid email and password, save user in database and send `201 - Created` response.

Request

```sh
curl \
  -X POST \
  http://localhost:1323/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"secret"}'
```
Response

`201 - Created`

```json
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "test@test.com",
  "password": "secret"
}
```


## Login
User login

Retrieves credentials from the body and validate against database.
For invalid credentials, send 401 - Unauthorized response.
For valid credentials, send 200 - OK response:
Generate JWT for the user and send it as response.
Each subsequent request must include JWT in the Authorization header.
Method: `POST`
Path: `/login`

Request

```sh
curl \
  -X POST \
  http://localhost:1323/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"secret"}'
```
Response

`200 - OK`

```json
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "test@test.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDI4NjAxMDUsImlkIjoiNWJmMjI4YTc1OTZhOTcyNDc0ODI5NTU1In0.UNN--g2Mn5y0JDFxo1XLmtu_rmvdTATSIKovVrC_vYo"
}
```
Client has now to implement token management.


## Create city
Creates city

For invalid request payload, send 400 - Bad Request response.
If user is not found, send 404 - Not Found response.
Otherwise save post in the database and return it via 201 - Created response.
Method: `POST`
Path: `/city/create`

Request

```sh
curl \
  -X POST \
  http://localhost:1323/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDI4NjAxMDUsImlkIjoiNWJmMjI4YTc1OTZhOTcyNDc0ODI5NTU1In0.UNN--g2Mn5y0JDFxo1XLmtu_rmvdTATSIKovVrC_vYo" \
  -H "Content-Type: application/json" \
  -d '{"name":"Singapore","desc":"Singapore Description","attractions":["Merlion","Marina Bay Sands"]}'
```
Response

`201 - Created`

```json
{
    "id": "5bf23a97596a972404f2de4a",
    "name": "Singapore",
    "desc": "Singapore Description",
    "attractions": [
        "Merlion",
        "Marina Bay Sands"
    ]
}
```

## Fetch cities
Fetches all cities

For invalid request payload, send 400 - Bad Request response.
If user is not found, send 404 - Not Found response.
Otherwise save post in the database and return it via 201 - Created response.
Method: `POST`
Path: `/city/all`

Request

```sh
curl \
  -X POST \
  http://localhost:1323/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDI4NjAxMDUsImlkIjoiNWJmMjI4YTc1OTZhOTcyNDc0ODI5NTU1In0.UNN--g2Mn5y0JDFxo1XLmtu_rmvdTATSIKovVrC_vYo" \
  -H "Content-Type: application/json" \
```
Response

`200 - OK`

```json
[
    {
        "id": "5bf23a97596a972404f2de4a",
        "name": "Singapore",
        "desc": "Singapore Description",
        "attractions": [
            "Merlion",
            "Marina Bay Sands"
        ]
    }
]
```

## Fetch city
Fetches specific city

For invalid request payload, send 400 - Bad Request response.
If user is not found, send 404 - Not Found response.
Otherwise save post in the database and return it via 201 - Created response.
Method: `POST`
Path: `/city/:name`

Request

```sh
curl \
  -X POST \
  http://localhost:1323/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDI4NjAxMDUsImlkIjoiNWJmMjI4YTc1OTZhOTcyNDc0ODI5NTU1In0.UNN--g2Mn5y0JDFxo1XLmtu_rmvdTATSIKovVrC_vYo" \
  -H "Content-Type: application/json" \
```
Response

`200 - OK`

```json
[
    {
        "id": "5bf23a97596a972404f2de4a",
        "name": "Singapore",
        "desc": "Singapore Description",
        "attractions": [
            "Merlion",
            "Marina Bay Sands"
        ]
    }
]
```

## License
[BSD 3-Clause](https://github.com/alvindcastro/alvindcastro/tree/master/travel-echo-mongo/LICENSE)

## TODO
Implement unit test (limited time)

Implement other logic (limited time)