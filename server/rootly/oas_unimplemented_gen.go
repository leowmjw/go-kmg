// Code generated by ogen, DO NOT EDIT.

package rootly

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// ScimUsersGet implements  operation.
//
// GET /scim/Users
func (UnimplementedHandler) ScimUsersGet(ctx context.Context, params ScimUsersGetParams) (r ScimUsersGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ScimUsersIDDelete implements  operation.
//
// DELETE /scim/Users/{id}
func (UnimplementedHandler) ScimUsersIDDelete(ctx context.Context, params ScimUsersIDDeleteParams) (r ScimUsersIDDeleteNoContent, _ error) {
	return r, ht.ErrNotImplemented
}

// ScimUsersIDGet implements  operation.
//
// GET /scim/Users/{id}
func (UnimplementedHandler) ScimUsersIDGet(ctx context.Context, params ScimUsersIDGetParams) (r ScimUsersIDGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ScimUsersIDPatch implements  operation.
//
// PATCH /scim/Users/{id}
func (UnimplementedHandler) ScimUsersIDPatch(ctx context.Context, params ScimUsersIDPatchParams) (r ScimUsersIDPatchOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ScimUsersPost implements  operation.
//
// POST /scim/Users
func (UnimplementedHandler) ScimUsersPost(ctx context.Context) (r ScimUsersPostCreated, _ error) {
	return r, ht.ErrNotImplemented
}
