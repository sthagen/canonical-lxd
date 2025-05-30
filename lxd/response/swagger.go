// Package response contains helpers for rendering LXD HTTP responses.
package response

import (
	"github.com/canonical/lxd/shared/api"
)

// Silence `unused` linter noise not realizing that these are used in the
// swagger docs (doc/rest-api.yaml).
var (
	_ = swaggerOperation{}
	_ = swaggerEmptySyncResponse{}
	_ = swaggerBadRequest{}
	_ = swaggerForbidden{}
	_ = swaggerPreconditionFailed{}
	_ = swaggerInternalServerError{}
	_ = swaggerNotFound{}
	_ = swaggerNotImplemented{}
	_ = swaggerEmptySyncResponse{}
)

// Operation
//
// swagger:response Operation
type swaggerOperation struct {
	// Empty sync response
	// in: body
	Body struct {
		// Example: async
		Type string `json:"type"`

		// Example: Operation created
		Status string `json:"status"`

		// Example: 100
		StatusCode int `json:"status_code"`

		// Example: /1.0/operations/66e83638-9dd7-4a26-aef2-5462814869a1
		Operation string `json:"operation"`

		Metadata api.Operation `json:"metadata"`
	}
}

// Empty sync response
//
// swagger:response EmptySyncResponse
type swaggerEmptySyncResponse struct {
	// Empty sync response
	// in: body
	Body struct {
		// Example: sync
		Type string `json:"type"`

		// Example: Success
		Status string `json:"status"`

		// Example: 200
		StatusCode int `json:"status_code"`
	}
}

// Bad Request
//
// swagger:response BadRequest
type swaggerBadRequest struct {
	// Bad Request
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: bad request
		Error string `json:"error"`

		// Example: 400
		ErrorCode int `json:"error_code"`
	}
}

// Forbidden
//
// swagger:response Forbidden
type swaggerForbidden struct {
	// Bad Request
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: not authorized
		Error string `json:"error"`

		// Example: 403
		ErrorCode int `json:"error_code"`
	}
}

// Precondition Failed
//
// swagger:response PreconditionFailed
type swaggerPreconditionFailed struct {
	// Internal server Error
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: precondition failed
		Error string `json:"error"`

		// Example: 412
		ErrorCode int `json:"error_code"`
	}
}

// Internal Server Error
//
// swagger:response InternalServerError
type swaggerInternalServerError struct {
	// Internal server Error
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: internal server error
		Error string `json:"error"`

		// Example: 500
		ErrorCode int `json:"error_code"`
	}
}

// Not found
//
// swagger:response NotFound
type swaggerNotFound struct {
	// Not found
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: not found
		Error string `json:"error"`

		// Example: 404
		ErrorCode int `json:"error_code"`
	}
}

// Not implemented
//
// swagger:response NotImplemented
type swaggerNotImplemented struct {
	// Not implemented
	// in: body
	Body struct {
		// Example: error
		Type string `json:"type"`

		// Example: not implemented
		Error string `json:"error"`

		// Example: 501
		ErrorCode int `json:"error_code"`
	}
}
