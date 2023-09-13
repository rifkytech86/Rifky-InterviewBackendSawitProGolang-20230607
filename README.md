# Backend Engineering Interview Assignment (Golang)

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.19
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

   Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

   Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project

To start working, execute

```
make init
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

To run test, run the following command:

```
make test
```


## Detail Update Code Test Interview Backend
Aplicant

| Name          | Email                 |
|---------------|-----------------------|
| Rifky Rachman |rifky.rachman@gmail.com|

### The Complete Project Folder Structure
```
📦 root
├─ bootstrap
│  ├─ app.go
│  ├─ bcrypt_hasher.go
│  ├─ env.go
│  ├─ jwt.go
│  ├─ logger.go
│  ├─ postgres.go
│  └─ validator.go
cmd
│  └─ main.go
├─ common
│  ├─ response_wapper.go
│  ├─ util.go
│  └─ validator.go
├─ errors
│  └─ errors.go
├─ generated
│  └─ api.gen.go
handler
│  ├─ endpoints.go
│  ├─ server.go
│  └─ user_service_handler.go
├─ models
│  └─ user_model.go
├─ repository
│  ├─ type.go
│  └─ user_service_repository
├─ vendor
├─ apy.yaml
├─ coverage.out
├─ database.sql
├─ docker-compose.yml
├─ Dockerfile
├─ go.mod
├─ Makefile
├─ private_key.pem
├─ public_key.pem
└─ README.md
└─ sawit-pro.postman_collection.json
```

### Explain Structure file
| Folder     | Description                                      | Example                                       |
|------------|--------------------------------------------------|-----------------------------------------------|
| bootstrap  | initial all depend with service                  | hasser, env, jwt, logger, database, validator |
| cmd        | main application for project                     |                                               |
| errors     | wrapper errors for this project                  |                                               |
| handler    | handler echo serve route                         |                                               |
| models     | for define field related with database           |                                               |
| repository | related connection or interaction with databases |                                               |

### Explain New Schema Table
| Field             | Data Type |       value |
|-------------------|:---------:|------------:|
| user_id           | serial    | PRIMARY KEY |
| user_phone_number | varchar   | 16 (uniqeu) |
| user_full_name    | varchar   |          60 |
| user_password     | varchar   |         200 |
| user_logged       | integer   |             |


### Example API Request and Response
- Registration
   - Request
```
 curl --location 'localhost:8080/registration' \
--header 'Content-Type: application/json' \
--data '{
    "phone_number": "+6223108782723",
    "password": "asdqweA!1",
    "full_name": "sawit-pro"
}'
```
- Response
```json
  {
    "code": 201,
    "message": "success",
    "data": {
        "user_id": 24
    }
}
```

- Login
   - Request
```
curl --location 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "phone_number": "+6223108782723",
    "password": "asdqweA!1"
}'
```
- Response
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "auth_jwt": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODU5ODUsInN1YiI6MjQsInRva2VuX3V1aWQiOiI0ZTY1OWI4YS1iODFhLTQ2YjktOGNhYS04NjIxYTBmNmU2ZjQifQ.3mZfYwFKAiP7had6zJu-TLZLXYvyjhPDDh2cJ_mHh71G232c4B_jsAvcBc5tUNsFLy-rRRNoBHC1s5_MSGUa9Q",
        "user_id": 24
    }
}
```

- Get Profile
   - Request
```
curl --location 'localhost:8080/get-profile' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODU5ODUsInN1YiI6MjQsInRva2VuX3V1aWQiOiI0ZTY1OWI4YS1iODFhLTQ2YjktOGNhYS04NjIxYTBmNmU2ZjQifQ.3mZfYwFKAiP7had6zJu-TLZLXYvyjhPDDh2cJ_mHh71G232c4B_jsAvcBc5tUNsFLy-rRRNoBHC1s5_MSGUa9Q'
```
- Response
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "full_name": "sawit-pro",
        "phone_number": "+6223108782723"
    }
}
```

- Update Profile
   - Request
```
curl --location --request PATCH 'localhost:8080/update-profile' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODU5ODUsInN1YiI6MjQsInRva2VuX3V1aWQiOiI0ZTY1OWI4YS1iODFhLTQ2YjktOGNhYS04NjIxYTBmNmU2ZjQifQ.3mZfYwFKAiP7had6zJu-TLZLXYvyjhPDDh2cJ_mHh71G232c4B_jsAvcBc5tUNsFLy-rRRNoBHC1s5_MSGUa9Q' \
--data '{
    "phone_number": "+6223108782723",
    "full_name": "sawit pro 2"
}'
```
- Response
```json
{
    "code": 202,
    "message": "success",
    "data": {
        "user_id": 22
    }
}
```

## include postman as well can download
- └─ sawit-pro.postman_collection.json

- Registration pre-script
```
function generateRandomPhoneNumber() {
    const countryCode = '+62';
    const randomDigits = Math.floor(Math.random() * 1e11).toString().padStart(11, '0');
    return countryCode + randomDigits;
}
const randomPhoneNumber = generateRandomPhoneNumber();
pm.collectionVariables.set("random-phone-number", randomPhoneNumber);
```

- Login test-script
```
var jsonData = pm.response.json();
pm.collectionVariables.set("bearer-token", jsonData.data.auth_jwt);
```
- Get Profile test-script
```
var jsonData = pm.response.json();
pm.collectionVariables.set("phone_number", jsonData.data.phone_number);
```










