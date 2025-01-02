package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// POST /users
	mux.HandleFunc("POST /users", registerUserHandler)

	// GET /users/{id}, DELETE /users/{id}
	mux.HandleFunc("GET /users/{id}", getUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	// POST /todos, GET /todos
	mux.HandleFunc("POST /todos", createTodoHandler)
	mux.HandleFunc("GET /todos", listTodosHandler)

	// GET /todos/{id}, PATCH /todos/{id}, DELETE /todos/{id}
	mux.HandleFunc("GET /todos/{id}", getTodoHandler)
	mux.HandleFunc("PATCH /todos/{id}", updateTodoHandler)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)

	// POST /todos/{id}/complete
	mux.HandleFunc("POST /todos/{id}/complete", completeTodoHandler)

	// POST /todos/{id}/reopen
	mux.HandleFunc("POST /todos/{id}/reopen", reopenTodoHandler)

	// PATCH /todos/{id}/priority
	mux.HandleFunc("PATCH /todos/{id}/priority", updateTodoPriorityHandler)

	// POST /todos/{id}/tags, DELETE /todos/{id}/tags
	mux.HandleFunc("POST /todos/{id}/tags", addTodoTagHandler)
	mux.HandleFunc("DELETE /todos/{id}/tags", removeTodoTagHandler)

	// GET /todos/filter
	mux.HandleFunc("GET /todos/filter", filterTodosHandler)

	return mux
}

// POST /users
func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.ID = 1 // Dummy ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GET /users/{id}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user := User{ID: intID, Name: "John Doe", Email: "john.doe@example.com", Role: "admin"} // Dummy data
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// POST /todos
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	todo.ID = 1 // Dummy ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GET /todos
func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{{ID: 1, Title: "Example", Description: "Example Description", Status: "UNFINISHED", Priority: "HIGH"}} // Dummy data
	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id}
func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid TODO ID", http.StatusBadRequest)
		return
	}
	todo := Todo{ID: intID, Title: "Example", Description: "Example Description", Status: "UNFINISHED", Priority: "HIGH"} // Dummy data
	json.NewEncoder(w).Encode(todo)
}

// PATCH /todos/{id}
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// DELETE /todos/{id}
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// GET /todos/filter
func filterTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Filtered TODOs")) // Placeholder response
}

// POST /todos/{id}/complete
func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO marked as completed")) // Placeholder response
}

// POST /todos/{id}/reopen
func reopenTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO reopened")) // Placeholder response
}

// PATCH /todos/{id}/priority
func updateTodoPriorityHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO priority updated")) // Placeholder response
}

// POST /todos/{id}/tags
func addTodoTagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tag added to TODO")) // Placeholder response
}

// DELETE /todos/{id}/tags
func removeTodoTagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
