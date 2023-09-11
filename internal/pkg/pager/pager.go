package pager

import "net/http"

var Default = New(DefaultLimit, DefaultOffset)

type Pager struct {
	limit  Limit
	offset Offset
}

// New is a constructor for Pager.
func New(limit Limit, offset Offset) Pager {
	return Pager{
		limit:  limit,
		offset: offset,
	}
}

// NewFromHTTPRequest is a constructor for Pager.
// It parses limit and offset from http request.
func NewFromHTTPRequest(req *http.Request) (Pager, error) {
	rawLimit := req.URL.Query().Get("limit")
	limit, err := NewLimitFromString(rawLimit)
	if err != nil {
		return Pager{}, err
	}

	rawOffset := req.URL.Query().Get("offset")
	offset, err := NewOffsetFromString(rawOffset)
	if err != nil {
		return Pager{}, err
	}

	return New(limit, offset), nil
}

func (p Pager) Limit() Limit   { return p.limit }
func (p Pager) Offset() Offset { return p.offset }
