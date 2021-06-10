// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetBookByIDParams creates a new GetBookByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBookByIDParams() *GetBookByIDParams {
	return &GetBookByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBookByIDParamsWithTimeout creates a new GetBookByIDParams object
// with the ability to set a timeout on a request.
func NewGetBookByIDParamsWithTimeout(timeout time.Duration) *GetBookByIDParams {
	return &GetBookByIDParams{
		timeout: timeout,
	}
}

// NewGetBookByIDParamsWithContext creates a new GetBookByIDParams object
// with the ability to set a context for a request.
func NewGetBookByIDParamsWithContext(ctx context.Context) *GetBookByIDParams {
	return &GetBookByIDParams{
		Context: ctx,
	}
}

// NewGetBookByIDParamsWithHTTPClient creates a new GetBookByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBookByIDParamsWithHTTPClient(client *http.Client) *GetBookByIDParams {
	return &GetBookByIDParams{
		HTTPClient: client,
	}
}

/* GetBookByIDParams contains all the parameters to send to the API endpoint
   for the get book by id operation.

   Typically these are written to a http.Request.
*/
type GetBookByIDParams struct {

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get book by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBookByIDParams) WithDefaults() *GetBookByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get book by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBookByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get book by id params
func (o *GetBookByIDParams) WithTimeout(timeout time.Duration) *GetBookByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get book by id params
func (o *GetBookByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get book by id params
func (o *GetBookByIDParams) WithContext(ctx context.Context) *GetBookByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get book by id params
func (o *GetBookByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get book by id params
func (o *GetBookByIDParams) WithHTTPClient(client *http.Client) *GetBookByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get book by id params
func (o *GetBookByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get book by id params
func (o *GetBookByIDParams) WithID(id int64) *GetBookByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get book by id params
func (o *GetBookByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBookByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
