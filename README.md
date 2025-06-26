
# 🧰 Features

📚 CRUD operations for books

    🔐 JWT-based authentication using HTTP Basic

    🧪 Token-protected endpoints

    🐳 Dockerized with a multi-stage build for efficiency

    ⚙️ CLI flags for port and auth toggle


# 📂 Folder Structure

    ├── Book_project.go         # Main application
    ├── Dockerfile              # Multi-stage Dockerfile
    ├── go.mod                  # Go modules
    ├── go.sum                  # Dependency checksums  


# 🚀 Getting Started

## ✅ 1. Clone the repository
    git clone git@github.com:Tutul-dhar/api_project.git
    cd api_project

## ✅ 2. Run the App (Locally)
    go run Book_project.go

    Default port: 8080
    Use --port=9090 to change it.

    Example:
        go run Book_project.go --port=9090 --auth=false

# 🔐 Authentication
### Get JWT token:
    curl -u AdminUser:AdminPassword http://localhost:8080/api/v1/get-token
### Use the token in Authorization header:
    curl -H "Authorization: BEARER <your_token>" http://localhost:8080/api/v1/books

# 📦 Docker Setup

## 🛠 1. Build the image
    docker build -t book_api:multi .

## ▶️ 2. Run the container
    docker run -d -p 8080:8080 book_api:multi

# 📡 API Endpoints

| Method    | Endpoint     | Description            |
| :-------- | :------- | :------------------------- |
| `pos`  | `/api/v1/books` | `Create a new book`|
| `GET`  | `/api/v1/books` | `List all books`|
| `GET`  | `/api/v1/books/{id}` | `Get a book by ID`|
| `PUT`  | `/api/v1/books/{id}` | `Update a book by ID`|
| `DELETE`  | `/api/v1/books/{id}` | `Delete a book by ID`|
| `GET`  | `/api/v1/get-token` | `Get JWT token (Basic)`|

# 💡 Usage Examples
### 🔐 Step 1: Get a JWT Token
    curl -u AdminUser:AdminPassword http://localhost:8080/api/v1/get-token

### Response:
    {
        "token": "<your_jwt_token_here>"
    }

## 📚 Book API Endpoints

### ✅ 1. Create a Book – POST /api/v1/books
    curl -X POST http://localhost:8080/api/v1/books \
    -H "Authorization: BEARER $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "The Go Programming Language",
        "authorList": ["Alan A. A. Donovan", "Brian W. Kernighan"],
        "publishDate": "2015-10-26",
        "isbn": "9780134190440"
    }'

### ✅ 2. List All Books – GET /api/v1/books
Replace <uuid> with the actual book UUID returned from create/list endpoints:

    curl http://localhost:8080/api/v1/books/<uuid> \
    -H "Authorization: BEARER $TOKEN"

### ✅ 3. Get a Single Book by ID – GET /api/v1/books/{id}

    curl http://localhost:8080/api/v1/books/<uuid> \
    -H "Authorization: BEARER $TOKEN"

### ✅ 4. Update a Book – PUT /api/v1/books/{id}

    curl -X PUT http://localhost:8080/api/v1/books/<uuid> \
    -H "Authorization: BEARER $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "The Go Programming Language (2nd Edition)",
        "authorList": ["Alan A. A. Donovan", "Brian W. Kernighan"],
        "publishDate": "2025-01-01",
        "isbn": "9780134190440"
    }'

### ✅ 5. Delete a Book – DELETE /api/v1/books/{id}

    curl -X DELETE http://localhost:8080/api/v1/books/<uuid> \
    -H "Authorization: BEARER $TOKEN"



