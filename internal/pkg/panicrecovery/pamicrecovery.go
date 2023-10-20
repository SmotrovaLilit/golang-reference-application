package panicrecovery

import (
	"context"
	"encoding/json"
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"log/slog"
	"net/http"
	xhttp "reference-application/internal/pkg/http"
	"reference-application/internal/pkg/log"
)

// HTTPHandlerMiddleware is a middleware that recovers from panics, and returns
func HTTPHandlerMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				ctx := kithttp.PopulateRequestContext(context.Background(), req)
				logger := log.WithContext(ctx, logger)
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				writer.WriteHeader(xhttp.ErrInternal.StatusCode)
				errEnc := json.NewEncoder(writer).Encode(xhttp.ErrInternal)
				if errEnc != nil {
					logger.Error(fmt.Sprintf("panic recovered and sending internal error in request is failed: (ecoding error: %s),  (panic error: %v)", errEnc.Error(), r))
				} else {
					logger.Error(fmt.Sprintf("panic recovered: %v", r))
				}
			}
		}()
		next.ServeHTTP(writer, req)
	})
}

type endpointType interface {
	~func(ctx context.Context, request interface{}) (response interface{}, err error)
}

// EndpointMiddleware is a middleware that recovers from panics, and returns
func EndpointMiddleware[T endpointType](e T, logger *slog.Logger) T {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger := log.WithContext(ctx, logger)
				logger.Error(fmt.Sprintf("panic recovered: %v", r))
				err = xhttp.ErrInternal
			}
		}()
		return e(ctx, request)
	}
}
