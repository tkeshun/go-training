package api

import (
	"context"
	"fmt"
	"pkg/logger"
	"todo/ogen"
)

type APIHandler struct {
	Logger *logger.Logger
}

// Ensure APIHandler implements the ogen.Handler interface.
var _ ogen.Handler = (*APIHandler)(nil)

// GET /todos/filter
func (h *APIHandler) TodosFilterGet(ctx context.Context, params ogen.TodosFilterGetParams) (ogen.TodosFilterGetRes, error) {
	return &ogen.TodosFilterGetOKApplicationJSON{}, nil
}

// GET /todos
func (h *APIHandler) TodosGet(ctx context.Context) ([]ogen.TodosGetOKItem, error) {
	return []ogen.TodosGetOKItem{}, nil
}

// POST /todos/{id}/complete
func (h *APIHandler) TodosIDCompletePost(ctx context.Context, params ogen.TodosIDCompletePostParams) (ogen.TodosIDCompletePostRes, error) {
	return &ogen.TodosIDCompletePostOK{}, nil
}

// TodosIDDelete implements DELETE /todos/{id} operation.
//
// Delete a TODO.
//
// DELETE /todos/{id}
func (h *APIHandler) TodosIDDelete(ctx context.Context, params ogen.TodosIDDeleteParams) (ogen.TodosIDDeleteRes, error) {
	return &ogen.TodosIDDeleteNoContent{}, nil
}

// TodosIDGet implements GET /todos/{id} operation.
//
// Get TODO details by ID.
//
// GET /todos/{id}
func (h *APIHandler) TodosIDGet(ctx context.Context, params ogen.TodosIDGetParams) (ogen.TodosIDGetRes, error) {
	return &ogen.TodosIDGetOK{}, nil
}

// TodosIDPatch implements PATCH /todos/{id} operation.
//
// Update a TODO.
//
// PATCH /todos/{id}
func (h *APIHandler) TodosIDPatch(ctx context.Context, req *ogen.TodosIDPatchReq, params ogen.TodosIDPatchParams) (ogen.TodosIDPatchRes, error) {
	return &ogen.TodosIDPatchOK{}, nil
}

// TodosIDPriorityPatch implements PATCH /todos/{id}/priority operation.
//
// Change the priority of a TODO.
//
// PATCH /todos/{id}/priority
func (h *APIHandler) TodosIDPriorityPatch(ctx context.Context, req *ogen.TodosIDPriorityPatchReq, params ogen.TodosIDPriorityPatchParams) (ogen.TodosIDPriorityPatchRes, error) {
	return &ogen.TodosIDPriorityPatchOK{}, nil
}

// TodosIDReopenPost implements POST /todos/{id}/reopen operation.
//
// Reopen a completed TODO.
//
// POST /todos/{id}/reopen
func (h *APIHandler) TodosIDReopenPost(ctx context.Context, params ogen.TodosIDReopenPostParams) (ogen.TodosIDReopenPostRes, error) {
	return &ogen.TodosIDReopenPostOK{}, nil
}

// TodosIDTagsDelete implements DELETE /todos/{id}/tags operation.
//
// Remove a tag from a TODO.
//
// DELETE /todos/{id}/tags
func (h *APIHandler) TodosIDTagsDelete(ctx context.Context, params ogen.TodosIDTagsDeleteParams) (ogen.TodosIDTagsDeleteRes, error) {
	return &ogen.TodosIDTagsDeleteNoContent{}, nil
}

// TodosIDTagsPost implements POST /todos/{id}/tags operation.
//
// Add a tag to a TODO.
//
// POST /todos/{id}/tags
func (h *APIHandler) TodosIDTagsPost(ctx context.Context, req *ogen.TodosIDTagsPostReq, params ogen.TodosIDTagsPostParams) (ogen.TodosIDTagsPostRes, error) {
	return &ogen.TodosIDTagsPostOK{}, nil
}

// TodosPost implements POST /todos operation.
//
// Create a new TODO.
//
// POST /todos
func (h *APIHandler) TodosPost(ctx context.Context, req *ogen.TodosPostReq) (ogen.TodosPostRes, error) {
	return &ogen.TodosPostCreated{}, nil
}

// UsersIDDelete implements DELETE /users/{id} operation.
//
// Delete a user and all related data.
//
// DELETE /users/{id}
func (h *APIHandler) UsersIDDelete(ctx context.Context, params ogen.UsersIDDeleteParams) (ogen.UsersIDDeleteRes, error) {
	// return &ogen.UsersIDDeleteNoContent{}, nil
	h.Logger.Info(ctx, fmt.Sprintf("access DELETE /users/{%v}", params.ID), nil)
	return &ogen.UsersIDDeleteNotFound{}, nil
}

// UsersIDGet implements GET /users/{id} operation.
//
// Get user details by ID.
//
// GET /users/{id}
func (h *APIHandler) UsersIDGet(ctx context.Context, params ogen.UsersIDGetParams) (ogen.UsersIDGetRes, error) {
	h.Logger.Info(ctx, fmt.Sprintf("access GET /users/{%v}", params.ID), nil)
	user := ogen.UsersIDGetOK{
		ID:    ogen.OptInt{Value: 1, Set: true},
		Name:  ogen.OptString{Value: "test", Set: true},
		Email: ogen.OptString{Value: "test", Set: true},
		Role:  ogen.OptString{Value: "test", Set: true},
	}

	return &user, nil
}

// UsersPost implements POST /users operation.
//
// Register a new user.
//
// POST /users
func (h *APIHandler) UsersPost(ctx context.Context, req *ogen.UsersPostReq) (ogen.UsersPostRes, error) {
	return &ogen.UsersPostCreated{}, nil
}
