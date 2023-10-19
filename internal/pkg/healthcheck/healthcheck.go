package healthcheck

import (
	"context"
	"github.com/alexliesenfeld/health"
	"net/http"
	"time"
)

// Checker preforms checks of components availability.
type Checker struct {
	options []health.CheckerOption
}

// CheckerFunc should return an error if the component is not available.
type CheckerFunc func(ctx context.Context) error

// New is a constructor.
func New(cacheDuration, timeout time.Duration) *Checker {
	opts := make([]health.CheckerOption, 0, 2)
	opts = append(opts, health.WithCacheDuration(cacheDuration))
	opts = append(opts, health.WithTimeout(timeout))
	return &Checker{
		options: opts,
	}
}

// Add adds checker to Checker.
func (h *Checker) Add(name string, check CheckerFunc) {
	h.options = append(h.options, health.WithCheck(health.Check{
		Name:  name,
		Check: check,
	}))
}

// Handler returns http.Handler that runs all checkers and writes result to response.
func (h *Checker) Handler() http.Handler {
	checker := health.NewChecker(h.options...)
	return health.NewHandler(checker)
}
