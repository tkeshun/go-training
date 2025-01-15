package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"todo/domain/model"
	"todo/domain/service"
	"todo/domain/value"

	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func SetupRoutes(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// POST /users
	mux.HandleFunc("POST /users", handler.registerUserHandler)

	// GET /users/{id}, DELETE /users/{id}
	mux.HandleFunc("GET /users/{id}", handler.getUserHandler)
	mux.HandleFunc("DELETE /users/{id}", handler.deleteUserHandler)

	// POST /todos, GET /todos
	mux.HandleFunc("POST /todos", handler.createTodoHandler)
	mux.HandleFunc("GET /todos", handler.listTodosHandler)

	// GET /todos/{id}, PATCH /todos/{id}, DELETE /todos/{id}
	mux.HandleFunc("GET /todos/{id}", handler.getTodoHandler)
	mux.HandleFunc("PATCH /todos/{id}", handler.updateTodoHandler)
	mux.HandleFunc("DELETE /todos/{id}", handler.deleteTodoHandler)

	// POST /todos/{id}/complete
	mux.HandleFunc("POST /todos/{id}/complete", handler.completeTodoHandler)

	// POST /todos/{id}/reopen
	mux.HandleFunc("POST /todos/{id}/reopen", handler.reopenTodoHandler)

	// PATCH /todos/{id}/priority
	mux.HandleFunc("PATCH /todos/{id}/priority", handler.updateTodoPriorityHandler)

	// POST /todos/{id}/tags, DELETE /todos/{id}/tags
	mux.HandleFunc("POST /todos/{id}/tags", handler.addTodoTagHandler)
	mux.HandleFunc("DELETE /todos/{id}/tags", handler.removeTodoTagHandler)

	// GET /todos/filter
	mux.HandleFunc("GET /todos/filter", handler.filterTodosHandler)

	return mux
}

type Handler struct {
	todoUseCase service.TodoService
	userUseCase service.UserService
}

func NewHandler(todo service.TodoService, user service.UserService) *Handler {
	return &Handler{
		todoUseCase: todo,
		userUseCase: user,
	}
}

// POST /users
func (h *Handler) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	pHash, err := hashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.userUseCase.RegisterUser(ctx, model.User{
		ID:    model.UserID(user.ID),
		Name:  model.UserName(user.Name),
		Email: model.UserEmail(user.Email),
		Role: func() value.Role {
			role, err := value.RoleString(user.Role)
			if err != nil {
				return value.REGULAR_USER
			}
			return role
		}(),
		PasswordHash: model.PasswordHash(pHash),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GET /users/{id}
func (h *Handler) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	user, err := h.userUseCase.GetUserByID(ctx, model.UserID(intID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func (h *Handler) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// POST /todos
func (h *Handler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
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
func (h *Handler) listTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{{ID: 1, Title: "Example", Description: "Example Description", Status: "UNFINISHED", Priority: "HIGH"}} // Dummy data
	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id}
func (h *Handler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
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
func (h *Handler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// DELETE /todos/{id}
func (h *Handler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// GET /todos/filter
func (h *Handler) filterTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Filtered TODOs")) // Placeholder response
}

// POST /todos/{id}/complete
func (h *Handler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO marked as completed")) // Placeholder response
}

// POST /todos/{id}/reopen
func (h *Handler) reopenTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO reopened")) // Placeholder response
}

// PATCH /todos/{id}/priority
func (h *Handler) updateTodoPriorityHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO priority updated")) // Placeholder response
}

// POST /todos/{id}/tags
func (h *Handler) addTodoTagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tag added to TODO")) // Placeholder response
}

// DELETE /todos/{id}/tags
func (h *Handler) removeTodoTagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
