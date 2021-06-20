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

	"github.com/MarlikAlmighty/library/pkg/models"
)

// NewAddBookParams creates a new AddBookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBookParams() *AddBookParams {
	return &AddBookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBookParamsWithTimeout creates a new AddBookParams object
// with the ability to set a timeout on a request.
func NewAddBookParamsWithTimeout(timeout time.Duration) *AddBookParams {
	return &AddBookParams{
		timeout: timeout,
	}
}

// NewAddBookParamsWithContext creates a new AddBookParams object
// with the ability to set a context for a request.
func NewAddBookParamsWithContext(ctx context.Context) *AddBookParams {
	return &AddBookParams{
		Context: ctx,
	}
}

// NewAddBookParamsWithHTTPClient creates a new AddBookParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBookParamsWithHTTPClient(client *http.Client) *AddBookParams {
	return &AddBookParams{
		HTTPClient: client,
	}
}

/* AddBookParams contains all the parameters to send to the API endpoint
   for the add book operation.

   Typically these are written to a http.Request.
*/
type AddBookParams struct {

	// Add.
	Add *models.Book

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBookParams) WithDefaults() *AddBookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add book params
func (o *AddBookParams) WithTimeout(timeout time.Duration) *AddBookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add book params
func (o *AddBookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add book params
func (o *AddBookParams) WithContext(ctx context.Context) *AddBookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add book params
func (o *AddBookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add book params
func (o *AddBookParams) WithHTTPClient(client *http.Client) *AddBookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add book params
func (o *AddBookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdd adds the add to the add book params
func (o *AddBookParams) WithAdd(add *models.Book) *AddBookParams {
	o.SetAdd(add)
	return o
}

// SetAdd adds the add to the add book params
func (o *AddBookParams) SetAdd(add *models.Book) {
	o.Add = add
}

// WriteToRequest writes these params to a swagger request
func (o *AddBookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Add != nil {
		if err := r.SetBodyParam(o.Add); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}