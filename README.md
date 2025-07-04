# Book Server API


## Prerequisites

- Go 1.24.3

- Docker

- Postman

- Git

## Setup and Run

### Clone the Repository:

    git clone https://github.com/Tutul-dhar/book-server.git
    cd book-server

### Install Dependencies:

    go mod tidy
    go mod download

### Run Locally:

    export JWT_SECRET="your-secret-key"
    go run -mod=mod ./cmd

### Build and Run with Docker:

    docker build -t book_project .
    docker run -e JWT_SECRET="your-secret-key" -p 8080:8080 book_project

### Run Tests:

    go test -mod=mod -v ./api/handler

## Postman Testing

### Register a User:

- Endpoint: POST {{base_url}}/api/v1/users/register

- Headers: Content-Type: application/json

- Body:

        {
            
            "username": "testuser",
            "password": "testpass"
        }

## Login:

- Endpoint: POST {{base_url}}/api/v1/users/login

- Headers: Content-Type: application/json

- Body:

        {
            "username": "testuser",
            "password": "testpass"
        } 

## Get JWT Token:

- Endpoint: GET {{base_url}}/api/v1/get-token

- Authorization: Basic Auth (Username: testuser, Password: testpass)

- Tests:

        const response = pm.response.json();
        pm.environment.set("jwt_token", response.token);

## List All Books:

- Endpoint: GET {{base_url}}/api/v1/books

- Authorization: Basic Auth (Username: testuser, Password: testpass)

## Create a Book:

- Endpoint: POST {{base_url}}/api/v1/books

- Headers:

        Content-Type: application/json

        Authorization: Bearer {{jwt_token}}

- Body:

        {
            "name": "The Go Programming Language",
            "authorList": ["Alan Donovan", "Brian Kernighan"],
            "publishDate": "2015-10-26",
            "isbn": "9780134190440"
        }

## Get a Book by ID:

- Endpoint: GET {{base_url}}/api/v1/books/<uuid>

- Headers: Authorization: Bearer {{jwt_token}}

## Update a Book:

- Endpoint: PUT {{base_url}}/api/v1/books/<uuid>

- Headers:

        Content-Type: application/json

        Authorization: Bearer {{jwt_token}}

- Body:

        {
            "name": "Updated Title",
            "authorList": ["New Author"],
            "publishDate": "2025-01-01",
            "isbn": "9780134190440"
        }

## Delete a Book:

- Endpoint: DELETE {{base_url}}/api/v1/books/<uuid>

- Headers: Authorization: Bearer {{jwt_token}}   

