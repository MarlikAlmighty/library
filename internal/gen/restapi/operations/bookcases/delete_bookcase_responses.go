// Code generated by go-swagger; DO NOT EDIT.

package bookcases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MarlikAlmighty/library/internal/gen/models"
)

// DeleteBookcaseOKCode is the HTTP code returned for type DeleteBookcaseOK
const DeleteBookcaseOKCode int = 200

/*DeleteBookcaseOK OK

swagger:response deleteBookcaseOK
*/
type DeleteBookcaseOK struct {
}

// NewDeleteBookcaseOK creates DeleteBookcaseOK with default headers values
func NewDeleteBookcaseOK() *DeleteBookcaseOK {

	return &DeleteBookcaseOK{}
}

// WriteResponse to the client
func (o *DeleteBookcaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteBookcaseBadRequestCode is the HTTP code returned for type DeleteBookcaseBadRequest
const DeleteBookcaseBadRequestCode int = 400

/*DeleteBookcaseBadRequest Bad request

swagger:response deleteBookcaseBadRequest
*/
type DeleteBookcaseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Fail `json:"body,omitempty"`
}

// NewDeleteBookcaseBadRequest creates DeleteBookcaseBadRequest with default headers values
func NewDeleteBookcaseBadRequest() *DeleteBookcaseBadRequest {

	return &DeleteBookcaseBadRequest{}
}

// WithPayload adds the payload to the delete bookcase bad request response
func (o *DeleteBookcaseBadRequest) WithPayload(payload *models.Fail) *DeleteBookcaseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bookcase bad request response
func (o *DeleteBookcaseBadRequest) SetPayload(payload *models.Fail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBookcaseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}