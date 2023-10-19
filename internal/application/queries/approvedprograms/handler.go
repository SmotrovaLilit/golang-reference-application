package approvedprograms

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"reference-application/internal/pkg/pager"
	"reference-application/internal/pkg/resource"
)

// Set is a wire set for programs query.
var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	NewEndpoint,
)

// Query is a query for programs.
type Query struct {
	pager pager.Pager
}

// NewQuery is a constructor for Query.
func NewQuery(pager pager.Pager) Query {
	return Query{pager: pager}
}

// Program is a result element of programs query.
type Program struct {
	ID           string
	PlatformCode string
	Version      Version
}

// Version is a part of programs query.
type Version struct {
	ID          string
	Name        string
	Number      string
	Description string
	Status      string
}

// Result is a result of programs query.
type Result []Program

// ReadModel is a read model for programs query.
// It is implemented by infrastructure layer.
type ReadModel interface {
	Query(ctx context.Context, pager pager.Pager) Result
}

// Handler is a handler for programs query.
type Handler struct {
	ReadModel ReadModel
}

// Handle handles a query for programs.
func (h Handler) Handle(ctx context.Context, query Query) Result {
	return h.ReadModel.Query(ctx, query.pager)
}

var _ resource.Endpoint = Endpoint(nil)

// Endpoint is an endpoint to update a version.
type Endpoint endpoint.Endpoint

// ResourceName returns the resource name.
// It uses for logging.
func (e Endpoint) ResourceName() string { return "program" }

// ResourceAction returns the resource action.
// It uses for logging.
func (e Endpoint) ResourceAction() string { return "get" }

// NewEndpoint creates a new endpoint for programs query.
func NewEndpoint(handler Handler) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(Query)
		return handler.Handle(ctx, query), nil
	}
}
