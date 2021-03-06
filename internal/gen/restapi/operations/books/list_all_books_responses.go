// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MarlikAlmighty/library/internal/gen/models"
)

// ListAllBooksOKCode is the HTTP code returned for type ListAllBooksOK
const ListAllBooksOKCode int = 200

/*ListAllBooksOK OK

swagger:response listAllBooksOK
*/
type ListAllBooksOK struct {

	/*
	  In: Body
	*/
	Payload models.Books `json:"body,omitempty"`
}

// NewListAllBooksOK creates ListAllBooksOK with default headers values
func NewListAllBooksOK() *ListAllBooksOK {

	return &ListAllBooksOK{}
}

// WithPayload adds the payload to the list all books o k response
func (o *ListAllBooksOK) WithPayload(payload models.Books) *ListAllBooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all books o k response
func (o *ListAllBooksOK) SetPayload(payload models.Books) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllBooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Books{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAllBooksBadRequestCode is the HTTP code returned for type ListAllBooksBadRequest
const ListAllBooksBadRequestCode int = 400

/*ListAllBooksBadRequest Bad request

swagger:response listAllBooksBadRequest
*/
type ListAllBooksBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Fail `json:"body,omitempty"`
}

// NewListAllBooksBadRequest creates ListAllBooksBadRequest with default headers values
func NewListAllBooksBadRequest() *ListAllBooksBadRequest {

	return &ListAllBooksBadRequest{}
}

// WithPayload adds the payload to the list all books bad request response
func (o *ListAllBooksBadRequest) WithPayload(payload *models.Fail) *ListAllBooksBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all books bad request response
func (o *ListAllBooksBadRequest) SetPayload(payload *models.Fail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllBooksBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
