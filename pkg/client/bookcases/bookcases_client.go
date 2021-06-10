// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bookcases API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bookcases API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddBookcase(params *AddBookcaseParams, opts ...ClientOption) (*AddBookcaseOK, error)

	DeleteBookcase(params *DeleteBookcaseParams, opts ...ClientOption) (*DeleteBookcaseOK, error)

	ListAllBookcases(params *ListAllBookcasesParams, opts ...ClientOption) (*ListAllBookcasesOK, error)

	ListAllBooksFromBookcases(params *ListAllBooksFromBookcasesParams, opts ...ClientOption) (*ListAllBooksFromBookcasesOK, error)

	PutBookcase(params *PutBookcaseParams, opts ...ClientOption) (*PutBookcaseOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddBookcase adds bookcase
*/
func (a *Client) AddBookcase(params *AddBookcaseParams, opts ...ClientOption) (*AddBookcaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBookcaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add_bookcase",
		Method:             "POST",
		PathPattern:        "/bookcases",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddBookcaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddBookcaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add_bookcase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBookcase deletes bookcase
*/
func (a *Client) DeleteBookcase(params *DeleteBookcaseParams, opts ...ClientOption) (*DeleteBookcaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBookcaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_bookcase",
		Method:             "DELETE",
		PathPattern:        "/bookcases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBookcaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBookcaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_bookcase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllBookcases shows all bookcases
*/
func (a *Client) ListAllBookcases(params *ListAllBookcasesParams, opts ...ClientOption) (*ListAllBookcasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllBookcasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list_all_bookcases",
		Method:             "GET",
		PathPattern:        "/bookcases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllBookcasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllBookcasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_all_bookcases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllBooksFromBookcases lists all books in bookcase
*/
func (a *Client) ListAllBooksFromBookcases(params *ListAllBooksFromBookcasesParams, opts ...ClientOption) (*ListAllBooksFromBookcasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllBooksFromBookcasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list_all_books_from_bookcases",
		Method:             "GET",
		PathPattern:        "/bookcases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllBooksFromBookcasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllBooksFromBookcasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_all_books_from_bookcases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutBookcase updates an existing bookcase
*/
func (a *Client) PutBookcase(params *PutBookcaseParams, opts ...ClientOption) (*PutBookcaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBookcaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_bookcase",
		Method:             "PUT",
		PathPattern:        "/bookcases",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBookcaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutBookcaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_bookcase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
