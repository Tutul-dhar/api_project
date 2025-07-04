package handler

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/tutul/book-server/api/middleware"
	"github.com/tutul/book-server/domain/entity"
	"github.com/tutul/book-server/infrastructure/persistence/inmemory"
	"github.com/tutul/book-server/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestServer(authEnabled bool) http.Handler {
	repo := inmemory.NewInMemoryRepo()
	svc := service.NewBookService(repo)
	bookHandler := NewBookHandler(svc)

	userRepo := inmemory.NewInMemoryUserRepo()
	userSvc := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userSvc)

	r := chi.NewRouter()
	RegisterRoutes(r, bookHandler, userHandler, middleware.TokenAuth, authEnabled)

	return r
}

func registerTestUser(t *testing.T, server http.Handler, username, password string) {
	user := entity.User{Username: username, Password: password}
	payload, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/api/v1/users/register", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req, server)
	if res.Code != http.StatusCreated {
		t.Fatalf("Failed to register user: %d", res.Code)
	}
}

func getToken(t *testing.T, server http.Handler, username, password string) string {
	req := httptest.NewRequest("GET", "/api/v1/get-token", nil)
	req.SetBasicAuth(username, password)
	res := executeRequest(req, server)
	if res.Code != http.StatusOK {
		t.Fatalf("Failed to get token: %d", res.Code)
	}
	var tokenResp map[string]string
	json.NewDecoder(res.Body).Decode(&tokenResp)
	return tokenResp["token"]
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected status %d, got %d", expected, actual)
	}
}

func Test_AllBookList(t *testing.T) {
	server := setupTestServer(true)

	// Register a test user
	username, password := "testuser", "testpass"
	registerTestUser(t, server, username, password)

	tests := []struct {
		method             string
		url                string
		body               io.Reader
		username           string
		password           string
		expectedStatusCode int
	}{
		{
			method:             "GET",
			url:                "/api/v1/books",
			body:               nil,
			username:           username,
			password:           password,
			expectedStatusCode: http.StatusOK,
		},
		{
			method:             "GET",
			url:                "/api/v1/books",
			body:               nil,
			username:           "wronguser",
			password:           password,
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range tests {
		req, _ := http.NewRequest(tc.method, tc.url, tc.body)
		req.SetBasicAuth(tc.username, tc.password)
		res := executeRequest(req, server)
		checkResponseCode(t, tc.expectedStatusCode, res.Code)
	}
}

func Test_CreateBook(t *testing.T) {
	server := setupTestServer(true)

	// Register a test user
	username, password := "testuser", "testpass"
	registerTestUser(t, server, username, password)

	// Get JWT token
	token := getToken(t, server, username, password)

	payload := `{
		"name": "The Go Programming Language",
		"authorList": ["Alan Donovan", "Brian Kernighan"],
		"publishDate": "2015-10-26",
		"isbn": "9780134190440"
	}`

	req, err := http.NewRequest("POST", "/api/v1/books", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res := executeRequest(req, server)

	checkResponseCode(t, http.StatusOK, res.Code)

	var createdBook map[string]interface{}
	if err := json.Unmarshal(res.Body.Bytes(), &createdBook); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if createdBook["name"] != "The Go Programming Language" {
		t.Errorf("Expected name to be %q, got %q", "The Go Programming Language", createdBook["name"])
	}
}

func Test_GetBook(t *testing.T) {
	server := setupTestServer(true)

	// Register a test user
	username, password := "testuser", "testpass"
	registerTestUser(t, server, username, password)

	// Get JWT token
	token := getToken(t, server, username, password)

	// Create book
	book := entity.Book{
		Name:        "Go Patterns",
		AuthorList:  []string{"Tutul"},
		PublishDate: "2024-01-01",
		ISBN:        "4444444444",
	}
	payload, _ := json.Marshal(book)
	createReq := httptest.NewRequest("POST", "/api/v1/books", bytes.NewReader(payload))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer "+token)
	createRes := executeRequest(createReq, server)

	var created entity.Book
	json.NewDecoder(createRes.Body).Decode(&created)

	// Get by ID
	getReq := httptest.NewRequest("GET", "/api/v1/books/"+created.UUID, nil)
	getReq.Header.Set("Authorization", "Bearer "+token)
	getRes := executeRequest(getReq, server)

	checkResponseCode(t, http.StatusOK, getRes.Code)
}

func Test_UpdateBook(t *testing.T) {
	server := setupTestServer(true)

	// Register a test user
	username, password := "testuser", "testpass"
	registerTestUser(t, server, username, password)

	// Get JWT token
	token := getToken(t, server, username, password)

	// Create book
	book := entity.Book{
		Name:        "Old Title",
		AuthorList:  []string{"Tutul"},
		PublishDate: "2020-01-01",
		ISBN:        "9999999999",
	}
	payload, _ := json.Marshal(book)
	createReq := httptest.NewRequest("POST", "/api/v1/books", bytes.NewReader(payload))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer "+token)
	createRes := executeRequest(createReq, server)

	var created entity.Book
	json.NewDecoder(createRes.Body).Decode(&created)

	// Update book
	updated := entity.Book{
		Name:        "New Title",
		AuthorList:  []string{"Author"},
		PublishDate: "2021-01-01",
		ISBN:        "9999999999",
	}
	updatePayload, _ := json.Marshal(updated)
	updateReq := httptest.NewRequest("PUT", "/api/v1/books/"+created.UUID, bytes.NewReader(updatePayload))
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set("Authorization", "Bearer "+token)
	updateRes := executeRequest(updateReq, server)

	checkResponseCode(t, http.StatusOK, updateRes.Code)

	var updatedBook entity.Book
	json.NewDecoder(updateRes.Body).Decode(&updatedBook)
	if updatedBook.Name != "New Title" {
		t.Errorf("Expected updated book name to be 'New Title', got %s", updatedBook.Name)
	}
}

func Test_DeleteBook(t *testing.T) {
	server := setupTestServer(true)

	// Register a test user
	username, password := "testuser", "testpass"
	registerTestUser(t, server, username, password)

	// Get JWT token
	token := getToken(t, server, username, password)

	// Create book
	book := entity.Book{
		Name:        "To be deleted",
		AuthorList:  []string{"Deleted person"},
		PublishDate: "2020-12-12",
		ISBN:        "1010101010",
	}
	payload, _ := json.Marshal(book)
	createReq := httptest.NewRequest("POST", "/api/v1/books", bytes.NewReader(payload))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer "+token)
	createRes := executeRequest(createReq, server)

	var created entity.Book
	json.NewDecoder(createRes.Body).Decode(&created)

	// Delete
	deleteReq := httptest.NewRequest("DELETE", "/api/v1/books/"+created.UUID, nil)
	deleteReq.Header.Set("Authorization", "Bearer "+token)
	deleteRes := executeRequest(deleteReq, server)

	checkResponseCode(t, http.StatusOK, deleteRes.Code)

	// Try to get again
	getReq := httptest.NewRequest("GET", "/api/v1/books/"+created.UUID, nil)
	getReq.Header.Set("Authorization", "Bearer "+token)
	getRes := executeRequest(getReq, server)

	checkResponseCode(t, http.StatusNotFound, getRes.Code)
}
