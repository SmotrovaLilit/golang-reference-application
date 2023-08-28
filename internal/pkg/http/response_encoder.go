package http

import (
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

// NoContentResponse is a response that contains no content.
type NoContentResponse struct{}

// StatusCode returns the status code for a no content response.
func (NoContentResponse) StatusCode() int { return http.StatusNoContent }

// NoContentResponseEncoder encodes a no content response.
// This is useful for endpoints that return no content.
func NoContentResponseEncoder(ctx context.Context, w http.ResponseWriter, _ interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, NoContentResponse{})
}
