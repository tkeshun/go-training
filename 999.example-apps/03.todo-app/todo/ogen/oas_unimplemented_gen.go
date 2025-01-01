// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// TodosFilterGet implements GET /todos/filter operation.
//
// Filter TODOs based on specific conditions.
//
// GET /todos/filter
func (UnimplementedHandler) TodosFilterGet(ctx context.Context, params TodosFilterGetParams) (r TodosFilterGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosGet implements GET /todos operation.
//
// Get a list of all TODOs.
//
// GET /todos
func (UnimplementedHandler) TodosGet(ctx context.Context) (r []TodosGetOKItem, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDCompletePost implements POST /todos/{id}/complete operation.
//
// Mark a TODO as completed.
//
// POST /todos/{id}/complete
func (UnimplementedHandler) TodosIDCompletePost(ctx context.Context, params TodosIDCompletePostParams) (r TodosIDCompletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDDelete implements DELETE /todos/{id} operation.
//
// Delete a TODO.
//
// DELETE /todos/{id}
func (UnimplementedHandler) TodosIDDelete(ctx context.Context, params TodosIDDeleteParams) (r TodosIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDGet implements GET /todos/{id} operation.
//
// Get TODO details by ID.
//
// GET /todos/{id}
func (UnimplementedHandler) TodosIDGet(ctx context.Context, params TodosIDGetParams) (r TodosIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDPatch implements PATCH /todos/{id} operation.
//
// Update a TODO.
//
// PATCH /todos/{id}
func (UnimplementedHandler) TodosIDPatch(ctx context.Context, req *TodosIDPatchReq, params TodosIDPatchParams) (r TodosIDPatchRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDPriorityPatch implements PATCH /todos/{id}/priority operation.
//
// Change the priority of a TODO.
//
// PATCH /todos/{id}/priority
func (UnimplementedHandler) TodosIDPriorityPatch(ctx context.Context, req *TodosIDPriorityPatchReq, params TodosIDPriorityPatchParams) (r TodosIDPriorityPatchRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDReopenPost implements POST /todos/{id}/reopen operation.
//
// Reopen a completed TODO.
//
// POST /todos/{id}/reopen
func (UnimplementedHandler) TodosIDReopenPost(ctx context.Context, params TodosIDReopenPostParams) (r TodosIDReopenPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDTagsDelete implements DELETE /todos/{id}/tags operation.
//
// Remove a tag from a TODO.
//
// DELETE /todos/{id}/tags
func (UnimplementedHandler) TodosIDTagsDelete(ctx context.Context, params TodosIDTagsDeleteParams) (r TodosIDTagsDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosIDTagsPost implements POST /todos/{id}/tags operation.
//
// Add a tag to a TODO.
//
// POST /todos/{id}/tags
func (UnimplementedHandler) TodosIDTagsPost(ctx context.Context, req *TodosIDTagsPostReq, params TodosIDTagsPostParams) (r TodosIDTagsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosPost implements POST /todos operation.
//
// Create a new TODO.
//
// POST /todos
func (UnimplementedHandler) TodosPost(ctx context.Context, req *TodosPostReq) (r TodosPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersIDDelete implements DELETE /users/{id} operation.
//
// Delete a user and all related data.
//
// DELETE /users/{id}
func (UnimplementedHandler) UsersIDDelete(ctx context.Context, params UsersIDDeleteParams) (r UsersIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersIDGet implements GET /users/{id} operation.
//
// Get user details by ID.
//
// GET /users/{id}
func (UnimplementedHandler) UsersIDGet(ctx context.Context, params UsersIDGetParams) (r UsersIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersPost implements POST /users operation.
//
// Register a new user.
//
// POST /users
func (UnimplementedHandler) UsersPost(ctx context.Context, req *UsersPostReq) (r UsersPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
