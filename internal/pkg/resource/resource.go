package resource

import (
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

type contextKey int

const (
	// ContextKeyResourceAction is the context key for the resource action.
	ContextKeyResourceAction contextKey = iota

	// ContextKeyResourceID is the context key for the resource ID.
	ContextKeyResourceID

	// ContextKeyResourceType is the context key for the resource type.
	ContextKeyResourceType
)

// Endpoint is an interface for endpoints that have a resource name and action.
type Endpoint interface {
	ResourceName() string
	ResourceAction() string
}

// PopulateContextRequestFunc returns a RequestFunc that populates the context with the resource ID and type.
func PopulateContextRequestFunc(resourceType, action string) kithttp.RequestFunc {
	return func(ctx context.Context, _ *http.Request) context.Context {
		ctx = context.WithValue(ctx, ContextKeyResourceAction, action)
		return context.WithValue(ctx, ContextKeyResourceType, resourceType)
	}
}

// PopulateContextWithResourceID populates the context with the resource ID.
func PopulateContextWithResourceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ContextKeyResourceID, id)
}
