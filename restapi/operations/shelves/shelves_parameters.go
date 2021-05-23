// Code generated by go-docs; DO NOT EDIT.

package shelves

// This file was generated by the docs tool.
// Editing this file might prove futile when you re-run the docs generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewShelvesParams creates a new ShelvesParams object
// no default values defined in spec.
func NewShelvesParams() ShelvesParams {

	return ShelvesParams{}
}

// ShelvesParams contains all the bound params for the shelves operation
// typically these are obtained from a http.Request
//
// docs:parameters shelves
type ShelvesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShelvesParams() beforehand.
func (o *ShelvesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
