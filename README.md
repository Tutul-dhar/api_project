Book Server API
This is a RESTful API built with Go using the Chi router, providing user management and book management functionalities. It includes endpoints for user registration, login, JWT token generation, and CRUD operations for books. The GET /api/v1/books endpoint requires Basic Authentication, while other book endpoints (POST, GET /{id}, PUT, DELETE) require JWT Authentication when enableAuth=true.
Project Structure
.
├── cmd
│   └── startProject.go
├── domain
│   ├── entity
│   │   ├── book.go
│   │   └── user.go
│   └── repository
│       ├── book_repository.go
│       └── user_repository.go
├── infrastructure
│   └── persistence
│       └── inmemory
│           ├── book.go
│           └── inmemory_user.go
├── service
│   ├── book_service.go
│   └── user_service.go
├── api
│   ├── handler
│   │   ├── book_handler.go
│   │   ├── routes.go
│   │   └── handler_test.go
│   └── middleware
│       ├── basicauth.go
│       ├── jwtauth.go
│       └── handler.go
├── go.mod
├── go.sum
└── Dockerfile

Prerequisites

Go: Version 1.24.3 or higher.
Docker: For building and running the containerized application.
Postman: For testing the API endpoints.
Git: To clone the repository.

Setup Instructions
Follow these steps to set up, run, and test the project.
1. Clone the Repository
Clone the project to your local machine:
git clone https://github.com/Tutul-dhar/book-server.git
cd book-server

2. Install Dependencies
Ensure you have Go 1.24.3 installed:
go version

Download the required dependencies specified in go.mod:
go mod tidy
go mod download

3. Run the Application Locally
Run the application using Go modules:
export JWT_SECRET="your-secret-key"
go run -mod=mod ./cmd


The server will start at http://localhost:8080.
The JWT_SECRET environment variable sets the JWT signing key (replace your-secret-key with a secure value).
To run without JWT Authentication for book endpoints (except GET /api/v1/books), use:go run -mod=mod ./cmd -auth=false



4. Build and Run with Docker
Build the Docker image:
docker build -t book_project .

Run the Docker container:
docker run -e JWT_SECRET="your-secret-key" -p 8080:8080 book_project


The server will be accessible at http://localhost:8080.
The -e JWT_SECRET sets the JWT secret for token generation.

5. Run Unit Tests
Run the unit tests to verify functionality:
go test -mod=mod -v ./api/handler

This tests user registration, login, token generation, and book CRUD operations, ensuring GET /api/v1/books uses Basic Authentication and other endpoints use JWT Authentication when enableAuth=true.
6. Test with Postman
Use Postman to test the API endpoints. Follow these steps to set up and test.
Postman Setup

Install Postman: Download from postman.com.
Create a Collection:
In Postman, click New > Collection, name it "Book Server API".


Set Up Environment:
Go to Environments (gear icon, top-right), create an environment named "Book Server".
Add variables:
base_url: http://localhost:8080
jwt_token: (leave blank; will be set after getting token)


Select the "Book Server" environment from the top-right dropdown.



API Endpoints
1. Register a User

Endpoint: POST {{base_url}}/api/v1/users/register
Headers:
Content-Type: application/json


Body (raw, JSON):{
  "username": "testuser",
  "password": "testpass"
}


Steps:
Create request: POST {{base_url}}/api/v1/users/register.
Headers: Add Content-Type: application/json.
Body: Select raw, JSON, paste the JSON.
Click Send.


Expected Response:
Status: 201 Created
Body:{
  "username": "testuser",
  "password": "[hashed_password]"
}





2. Login

Endpoint: POST {{base_url}}/api/v1/users/login
Headers:
Content-Type: application/json


Body (raw, JSON):{
  "username": "testuser",
  "password": "testpass"
}


Steps:
Create request: POST {{base_url}}/api/v1/users/login.
Headers: Add Content-Type: application/json.
Body: Select raw, JSON, paste the JSON.
Click Send.


Expected Response:
Status: 200 OK
Body:{
  "username": "testuser",
  "password": "[hashed_password]"
}





3. Get JWT Token

Endpoint: GET {{base_url}}/api/v1/get-token
Authentication: Basic Auth
Username: testuser
Password: testpass


Steps:
Create request: GET {{base_url}}/api/v1/get-token.
Authorization: Select Basic Auth, enter Username: testuser, Password: testpass.
Tests: Add script to save token:const response = pm.response.json();
pm.environment.set("jwt_token", response.token);


Click Send.


Expected Response:
Status: 200 OK
Body:{
  "token": "<jwt_token>"
}





4. List All Books

Endpoint: GET {{base_url}}/api/v1/books
Authentication: Basic Auth
Username: testuser
Password: testpass


Steps:
Create request: GET {{base_url}}/api/v1/books.
Authorization: Select Basic Auth, enter Username: testuser, Password: testpass.
Click Send.


Expected Response:
Status: 200 OK
Body: [] (empty array if no books exist)



5. Create a Book

Endpoint: POST {{base_url}}/api/v1/books
Headers:
Content-Type: application/json
Authorization: Bearer {{jwt_token}}


Body (raw, JSON):{
  "name": "The Go Programming Language",
  "authorList": ["Alan Donovan", "Brian Kernighan"],
  "publishDate": "2015-10-26",
  "isbn": "9780134190440"
}


Steps:
Create request: POST {{base_url}}/api/v1/books.
Headers:
Content-Type: application/json
Authorization: Bearer {{jwt_token}}


Body: Select raw, JSON, paste the JSON.
Click Send.


Expected Response:
Status: 200 OK
Body:{
  "uuid": "<generated_uuid>",
  "name": "The Go Programming Language",
  "authorList": ["Alan Donovan", "Brian Kernighan"],
  "publishDate": "2015-10-26",
  "isbn": "9780134190440"
}





6. Get a Book by ID

Endpoint: GET {{base_url}}/api/v1/books/<uuid>
Headers:
Authorization: Bearer {{jwt_token}}


Steps:
Create request: GET {{base_url}}/api/v1/books/<uuid> (replace <uuid> with the UUID from the create response).
Headers: Add Authorization: Bearer {{jwt_token}}.
Click Send.


Expected Response:
Status: 200 OK
Body: Book details.



7. Update a Book

Endpoint: PUT {{base_url}}/api/v1/books/<uuid>
Headers:
Content-Type: application/json
Authorization: Bearer {{jwt_token}}


Body (raw, JSON):{
  "name": "Updated Title",
  "authorList": ["New Author"],
  "publishDate": "2025-01-01",
  "isbn": "9780134190440"
}


Steps:
Create request: PUT {{base_url}}/api/v1/books/<uuid>.
Headers:
Content-Type: application/json
Authorization: Bearer {{jwt_token}}


Body: Select raw, JSON, paste the JSON.
Click Send.


Expected Response:
Status: 200 OK
Body: Updated book details.



8. Delete a Book

Endpoint: DELETE {{base_url}}/api/v1/books/<uuid>
Headers:
Authorization: Bearer {{jwt_token}}


Steps:
Create request: DELETE {{base_url}}/api/v1/books/<uuid>.
Headers: Add Authorization: Bearer {{jwt_token}}.
Click Send.


Expected Response:
Status: 200 OK
Body: Deleted book details.



API Flow

Register a user (POST /api/v1/users/register) to create credentials.
Login (POST /api/v1/users/login) to verify credentials (optional).
Get a JWT token (GET /api/v1/get-token) using Basic Authentication (username/password).
List books (GET /api/v1/books) using Basic Authentication to retrieve all books.
Manage books:
Create (POST /api/v1/books), get (GET /api/v1/books/<uuid>), update (PUT /api/v1/books/<uuid>), or delete (DELETE /api/v1/books/<uuid>) books using JWT Authentication (Bearer token).



Troubleshooting

Build Errors:
If docker build fails, ensure the Go version in Dockerfile (golang:1.24.3) matches go.mod (go 1.24.3).
Clear Docker cache:docker builder prune
docker build -t book_project .




Test Failures:
Run go test -mod=mod -v ./api/handler and check the output for errors.
Ensure api/handler/handler_test.go and api/handler/routes.go match the provided versions.


Postman Errors:
401 Unauthorized for GET /api/v1/books: Verify Basic Auth credentials (testuser:testpass).
401 Unauthorized for other book endpoints: Ensure the jwt_token variable is set correctly.
405 Method Not Allowed: Use GET for /api/v1/get-token, not POST.
