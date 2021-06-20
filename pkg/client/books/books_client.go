// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new books API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for books API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddBook(params *AddBookParams, opts ...ClientOption) (*AddBookOK, error)

	DeleteBook(params *DeleteBookParams, opts ...ClientOption) (*DeleteBookOK, error)

	GetBookByID(params *GetBookByIDParams, opts ...ClientOption) (*GetBookByIDOK, error)

	ListAllBooks(params *ListAllBooksParams, opts ...ClientOption) (*ListAllBooksOK, error)

	PutBook(params *PutBookParams, opts ...ClientOption) (*PutBookOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddBook adds book
*/
func (a *Client) AddBook(params *AddBookParams, opts ...ClientOption) (*AddBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add_book",
		Method:             "POST",
		PathPattern:        "/books",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddBookReader{formats: a.formats},
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
	success, ok := result.(*AddBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add_book: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBook deletes book
*/
func (a *Client) DeleteBook(params *DeleteBookParams, opts ...ClientOption) (*DeleteBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_book",
		Method:             "DELETE",
		PathPattern:        "/book/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBookReader{formats: a.formats},
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
	success, ok := result.(*DeleteBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_book: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBookByID gets book by id
*/
func (a *Client) GetBookByID(params *GetBookByIDParams, opts ...ClientOption) (*GetBookByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBookByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_book_by_id",
		Method:             "GET",
		PathPattern:        "/book/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBookByIDReader{formats: a.formats},
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
	success, ok := result.(*GetBookByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_book_by_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllBooks shows all books
*/
func (a *Client) ListAllBooks(params *ListAllBooksParams, opts ...ClientOption) (*ListAllBooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllBooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list_all_books",
		Method:             "GET",
		PathPattern:        "/books",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllBooksReader{formats: a.formats},
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
	success, ok := result.(*ListAllBooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_all_books: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutBook updates an existing book
*/
func (a *Client) PutBook(params *PutBookParams, opts ...ClientOption) (*PutBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_book",
		Method:             "PUT",
		PathPattern:        "/books",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBookReader{formats: a.formats},
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
	success, ok := result.(*PutBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_book: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}